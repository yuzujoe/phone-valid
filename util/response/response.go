package response

import "errors"

var (
	// InternalServerError 500 error response message
	InternalServerError = serverErrorMessage()
	// UserSignup400Reponse user/signin api 400 response messege
	UserSignup400Reponse = signup400Response()
	// UserAuthentication400Response  user/authentication api 400 response messege
	UserAuthentication400Response = "Your input code not correct"
	// UserAuthentication401Response user/authentication api 401 response messege
	UserAuthentication401Response = "AuthCode Expired, please get the authorization code again and enter"
	// UserAuthentication404Response user/authentication api 404 response messege
	UserAuthentication404Response = "User not exists"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UserAuthenticationSuccessReponse struct {
	Code  int64  `json:"code"`
	Token string `json:"token"`
}

func signup400Response() error { return errors.New("Please enter the correct phone number") }

func serverErrorMessage() error {
	return errors.New("sorry, server error occurring, please wait for recovery")
}

func SignupBadRequestResponse() error {
	return UserSignup400Reponse
}

func ServerErrorResponse() error {
	return InternalServerError
}
