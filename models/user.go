package models

import "time"

type User struct {
	UserID int64 `gorm:"primary_key; AUTO_INCREMENT" json:"user_id" form:"user_id"`

	PhoneNumber string `gorm:"unique;not null" json:"phone_number" form:"phone_number"`

	Code *AuthenticationCode `gorm:"foreignkey:PhoneNumber"`

	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}
