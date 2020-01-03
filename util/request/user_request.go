package request

type UserSignupRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UserAuthenticationRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
}
