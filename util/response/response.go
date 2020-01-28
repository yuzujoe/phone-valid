package response

import "errors"

var (
	// InternalServerError 500 error response message
	InternalServerError = serverErrorMessage()
	// UserSignup400Reponse user/signin api 400 response messege
	UserSignup400Reponse = signup400Response()
	// UserAuthentication400Response  user/authentication api 400 response messege
	UserAuthentication400Response = auth400Response()
	// UserAuthentication403Response user/authentication api 403 response messege
	UserAuthentication403Response = auth403Response()
	// UserAuthentication404Response user/authentication api 404 response messege
	UserAuthentication404Response = auth404Response()
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

func auth400Response() error {
	return errors.New("Your input code not correct")
}

func auth403Response() error {
	return errors.New("AuthCode Expired, please get the authorization code again and enter")
}

func auth404Response() error {
	return errors.New("User not exists")
}

func SignupBadRequestResponse() error {
	return UserSignup400Reponse
}

func ServerErrorResponse() error {
	return InternalServerError
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
