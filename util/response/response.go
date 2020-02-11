package response

import "errors"

var (
	// ErrInternalServerError 500 error response message
	ErrInternalServerError = errors.New("sorry, server error occurring, please wait for recovery")
	// UserSignup400Reponse user/signin api 400 response messege
	UserSignup400Reponse = errors.New("Please enter the correct phone number")
	// UserAuthentication400Response  user/authentication api 400 response messege
	UserAuthentication400Response = errors.New("Your input code not correct")
	// UserAuthentication403Response user/authentication api 403 response messege
	UserAuthentication403Response = errors.New("AuthCode Expired, please get the authorization code again and enter")
	// UserAuthentication404Response user/authentication api 404 response messege
	UserAuthentication404Response = errors.New("User not exists")
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type RequestToken struct {
	Code         int64  `json:"code"`
	RequestToken string `json:"request_token"`
	Message      string `json:"message"`
}

type UserAuthenticationSuccessReponse struct {
	Code  int64  `json:"code"`
	Token string `json:"token"`
}

func SignupBadRequestResponse() error {
	return UserSignup400Reponse
}

func ServerErrorResponse() error {
	return ErrInternalServerError
}

func Authenticate400Err() error {
	return UserAuthentication400Response
}

func Auth403Err() error {
	return UserAuthentication403Response
}

func Auth404Err() error {
	return UserAuthentication404Response
}
