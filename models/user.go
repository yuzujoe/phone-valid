package models

import "time"

type User struct {
	UserID int64 `gorm:"primary_key; AUTO_INCREMENT" json:"user_id" form:"user_id"`

	FirstName string `gorm:"first_name" json:"first_name" form:"first_name"`

	LastName string `gorm:"last_name" json:"last_name" form:"last_name"`

	PhoneNumber string `gorm:"unique;not null" json:"phone_number" form:"phone_number"`

	Gender string `gorm:"type: enum('male', 'female', 'unspecified'); default: 'unspecified';" json:"gender" form:"gender"`

	Birthday *time.Time `gorm:"birthday" json:"birthday" form:"birthday"`

	Zipcode string `gorm:"zipcode" json:"zipcode" form:"zipcode"`

	Prefecture string `gorm:"column:prefecture" json:"prefecture" form:"prefecture"`

	City string `gorm:"column:city" json:"city" form:"city"`

	AddressLine string `gorm:"column:address_line" json:"address_line" form:"address_line"`

	Code string `gorm:"column:code; NOT NULL" json:"code" form:"code"`

	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}
