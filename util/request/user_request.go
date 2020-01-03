package request

import "time"

type UserSignupRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UserAuthenticationRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

type CreateProfileRequest struct {
	Email       string    `json:"email" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	Birthday    time.Time `json:"birthday" binding:"required"`
	Zipcode     string    `json:"zipcode" binding:"required"`
	Prefecture  string    `json:"prefecture" binding:"required"`
	City        string    `json:"city" binding:"required"`
	AddressLine string    `json:"address_line" binding:"required"`
}
