package models

import "time"

type User struct {
	UserID int64 `gorm:"primary_key" json:"user_id" form:"user_id"`

	Email string `gorm:"email" json:"email" form:"email"`

	FirstName string `gorm:"firstName" json:"firstName" form:"firstName"`

	LastName string `gorm:"lastName" json:"lastName" form:"lastName"`

	PhoneNumber string `gorm:"unique;not null" json:"phone_number" form:"phone_number"`

	Gender string `gorm:"type: enum('male', 'female');" json:"gender" form:"gender"`

	Birthday time.Time `gorm:"birthday; type:date" json:"birthday" form:"birthday"`

	Zipcode string `gorm:"zipcode" json:"zipcode" form:"zipcode"`

	Prefecture string `gorm:"column:prefecture" json:"prefecture" form:"prefecture"`

	City string `gorm:"column:city" json:"city" form:"city"`

	AddressLine string `gorm:"column:address_line" json:"address_line" form:"address_line"`

	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}
