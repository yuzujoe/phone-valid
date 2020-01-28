package controller

import (
	"net/http"
	"phone-valid/util/auth"
	"phone-valid/util/jwt"
	"phone-valid/util/request"
	"phone-valid/util/response"
	"phone-valid/util/sms"
	"strconv"
	"time"

	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const codeLength = 6

func userSignupImpl(c *gin.Context, request *request.UserSignupRequest) (*response.Response, error) {
	var err error

	phoneValid := phoneValid(request.PhoneNumber)
	if !phoneValid {
		err = response.SignupBadRequestResponse()
		return nil, err
	}

	if err := createUser(request.PhoneNumber); err != nil {
		err = response.ServerErrorResponse()
		return nil, err
	}

	code := auth.GenerateAuthCode(codeLength)
	if err := registerCode(request.PhoneNumber, code); err != nil {
		err = response.ServerErrorResponse()
		return nil, err
	}

	if err := sms.PushSms(request.PhoneNumber, code); err != nil {
		err = response.ServerErrorResponse()
		return nil, err
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: "A 6-digit confirmation code has been sent to the entered phone number",
	}, nil
}

func userAuthenticationImpl(c *gin.Context, request *request.UserAuthenticationRequest) (*response.UserAuthenticationSuccessReponse, error) {

	authCode, err := getCodeInfo(request.PhoneNumber)
	if err != nil {
		err = response.Authenticate400Err()
		return nil, err
	}

	if err := compareCode(authCode.Code, request.Code); err != nil {
		if err != nil {
			err = response.Authenticate400Err()
			return nil, err
		}
	}

	if err := checkExpired(authCode.Expired); err != nil {
		err = response.Auth403Err()
		return nil, err
	}

	user, err := userExist(request.PhoneNumber)
	if err != nil {
		err = response.Auth404Err()
		return nil, err
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
		err = response.ServerErrorResponse()
		return nil, err
	}

	return &response.Response{
		Code:    200,
		Message: "UserCreateProfileSuccessResponse",
	}, nil
}
