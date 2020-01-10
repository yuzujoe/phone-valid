package controller

import (
	"log"
	"net/http"
	"phone-valid/models"
	"phone-valid/mysql"
	"phone-valid/util/auth"
	"phone-valid/util/request"
	"phone-valid/util/response"
	"phone-valid/util/sms"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

const codeLength = 6

func userSignupImpl(c *gin.Context, request *request.UserSignupRequest) (*response.Response, error) {

	if err := phoneValid(request.PhoneNumber); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    http.StatusBadRequest,
			Message: "The value entered is incorrect phone number format",
		})
		return nil, err
	}

	if err := createUser(request.PhoneNumber); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Connection to server failed, please try again in a good communication environment",
		})
		return nil, err
	}

	code := auth.GenerateAuthCode(codeLength)
	if err := registerCode(request.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Connection to server failed, please try again in a good communication environment",
		})
		return nil, err
	}

	if err := sms.PushSms(request.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "通信に失敗しました",
		})
		return nil, err
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: "A 6-digit confirmation code has been sent to the entered phone number",
	}, nil
}

func phoneValid(phoneNumber string) error {
	policy := "^\\d{2,4}-?\\d{2,4}-?\\d{3,4}$"
	re := regexp.MustCompile(policy)
	reg := re.MatchString(phoneNumber)
	if !reg {
		return nil
	}
	return nil
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

func userExist(phoneNumber string) *models.User {
	db := mysql.DB

	var user models.User

	auth := db.Where("phone_number = ?", phoneNumber).Select("user_id, phone_number").First(&user).RecordNotFound()

	if auth {
		return nil
	}

	return &user
}

func getCodeInfo(phoneNumber string) *models.AuthenticationCode {
	db := mysql.DB

	var authentication models.AuthenticationCode

	authCode := db.Where("phone_number = ?", phoneNumber).Select("phone_number, code, expired").First(&authentication).RecordNotFound()

	if authCode {
		return nil
	}

	return &authentication
}

func compareCode(code, reqCode string) bool {
	if code != reqCode {
		return false
	}
	return true
}

func checkExpired(expired time.Time) (bool, error) {
	old := expired
	now := time.Now()

	if old.Before(now) {
		return false, nil
	}

	return true, nil
}
