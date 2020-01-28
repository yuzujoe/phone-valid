package models

import "time"

// AuthenticationCode models authentication_codes
type AuthenticationCode struct {
	Code string `gorm:"column:code; not null" json:"code" form:"code"`

	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`

	RequestToken string `gorm:"column:request_token" json:"request_token"`

	Expired time.Time `gorm:"colunm:expired; not null" json:"expired" form:"expired"`

	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}
