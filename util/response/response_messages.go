package response

const (
	// InternalServerError 500 error response message
	InternalServerError = "sorry, server error occurring, please wait for recovery"
	// UserSignup400Reponse user/signin api 400 response messege
	UserSignup400Reponse = "Please enter the correct phone number"
	// UserAuthentication400Response  user/authentication api 400 response messege
	UserAuthentication400Response = "Your input code not correct"
	// UserAuthentication401Response user/authentication api 401 response messege
	UserAuthentication401Response = "AuthCode Expired, please get the authorization code again and enter"
	// UserAuthentication404Response user/authentication api 404 response messege
	UserAuthentication404Response = "User not exists"
	// UserCreateProfileSuccessResponse post user/profile api 200 response
	UserCreateProfileSuccessResponse = "Created Profile"
)
