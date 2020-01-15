package controller

import (
	"fmt"
	"log"
	"net/http"
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
			Code:    400,
			Message: response.UserSignup400Reponse,
		})
		return &response.Response{}
	}

	if err := createUser(request.PhoneNumber); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    500,
			Message: response.InternalServerError,
		})
		return &response.Response{}
	}

	code := auth.GenerateAuthCode(codeLength)
	if err := registerCode(request.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    500,
			Message: response.InternalServerError,
		})
		return &response.Response{}
	}

	if err := sms.PushSms(request.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:    500,
			Message: response.InternalServerError,
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
			Message: response.UserAuthentication400Response,
		})
		return nil, err
	}

	if err := compareCode(authCode.Code, request.Code); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    400,
			Message: response.UserAuthentication400Response,
		})
		return nil, err
	}

	if err := checkExpired(authCode.Expired); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code:    401,
			Message: response.UserAuthentication401Response,
		})
		return nil, err
	}
	fmt.Println("ok")
	user, err := userExist(request.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, &response.Response{
			Code:    404,
			Message: response.UserAuthentication404Response,
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
			Message: response.InternalServerError,
		})
		return nil, err
	}

	return &response.Response{
		Code:    200,
		Message: response.UserCreateProfileSuccessResponse,
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
