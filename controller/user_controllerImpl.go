package controller

import (
	"fmt"
	"log"
	"net/http"
	"phone-valid/models"
	"phone-valid/mysql"
	"phone-valid/util/auth"
	"phone-valid/util/jwt"
	"phone-valid/util/request"
	"phone-valid/util/response"
	"phone-valid/util/sms"
	"regexp"
	"strconv"
	"time"

	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const codeLength = 6

func userSignupImpl(c *gin.Context, request *request.UserSignupRequest) *response.Response {

	phoneLegCheck := phoneValid(request.PhoneNumber)
	if !phoneLegCheck {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    http.StatusBadRequest,
			Message: "The value entered is incorrect phone number format",
		})
		return &response.Response{}
	}

	if err := createUser(request.PhoneNumber); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Connection to server failed, please try again in a good communication environment",
		})
		return &response.Response{}
	}

	code := auth.GenerateAuthCode(codeLength)
	if err := registerCode(request.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Connection to server failed, please try again in a good communication environment",
		})
		return &response.Response{}
	}

	if err := sms.PushSms(request.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "通信に失敗しました",
		})
		return &response.Response{}
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: "A 6-digit confirmation code has been sent to the entered phone number",
	}
}

func userAuthenticationImpl(c *gin.Context, request *request.UserAuthenticationRequest) (*response.UserAuthenticationSuccessReponse, error) {

	authCode, err := getCodeInfo(request.PhoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    400,
			Message: "your input code not correct",
		})
		return nil, err
	}

	if err := compareCode(authCode.Code, request.Code); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    400,
			Message: "your input code not correct",
		})
		return nil, err
	}

	if err := checkExpired(authCode.Expired); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    401,
			Message: "Expired, please get the authorization code again and enter",
		})
		return nil, err
	}
	fmt.Println("ok")
	user, err := userExist(request.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, &response.Response{
			Code:    404,
			Message: "user not exists",
		})
		return nil, nil
	}

	userID := strconv.FormatInt(int64(user.UserID), 10)

	token, _ := jwt.TokenGenerate(jwt_go.MapClaims{
		"UserID": userID,
		"expire": time.Now().Add(time.Hour * 2).Unix(),
	})

	return &response.UserAuthenticationSuccessReponse{
		Code:  200,
		Token: token,
	}, nil
}

func userProfileCreateImpl(c *gin.Context, request *request.CreateProfileRequest) (*response.Response, error) {
	if err := insertProfile(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:    500,
			Message: err.Error(),
		})
		return nil, err
	}

	return &response.Response{
		Code:    200,
		Message: "Created Profile",
	}, nil
}

func phoneValid(phoneNumber string) bool {
	policy := "^\\d{2,4}-?\\d{2,4}-?\\d{3,4}$"
	re := regexp.MustCompile(policy)
	reg := re.MatchString(phoneNumber)
	if !reg {
		return false
	}
	return true
}

func createUser(phoneNumber string) error {

	db := mysql.DB

	var user models.User
	if err := db.FirstOrCreate(&user, models.User{PhoneNumber: phoneNumber}).Error; err != nil {
		log.Fatalln(err)
		return nil
	}

	return nil
}

func registerCode(phoneNumber, code string) error {
	db := mysql.DB

	var authCode models.AuthenticationCode

	expired := time.Now().Add(15 * time.Minute)

	if err := db.Where(models.AuthenticationCode{PhoneNumber: phoneNumber}).Assign(models.AuthenticationCode{Code: code, Expired: expired, UpdatedAt: time.Now()}).FirstOrCreate(&authCode).Error; err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func userExist(phoneNumber string) (models.User, error) {
	db := mysql.DB

	var user models.User

	if err := db.Where("phone_number = ?", phoneNumber).Select("user_id, phone_number").First(&user).Error; err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func getCodeInfo(phoneNumber string) (*models.AuthenticationCode, error) {
	db := mysql.DB

	var authentication models.AuthenticationCode

	authCode := db.Where("phone_number = ?", phoneNumber).Select("phone_number, code, expired").First(&authentication).RecordNotFound()

	if authCode {
		return nil, db.Error
	}

	return &authentication, nil
}

func compareCode(code, reqCode string) error {
	if code != reqCode {
		return http.ErrAbortHandler
	}
	return nil
}

func checkExpired(expired time.Time) error {
	old := expired
	now := time.Now()

	if old.Before(now) {
		return http.ErrAbortHandler
	}

	return nil
}
