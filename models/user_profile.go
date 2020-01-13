package models

import "time"

type UserProfile struct {
	UserID      int64     `json:"user_id" form:"user_id"`
	Email       string    `gorm:"not null" json:"email"`
	FirstName   string    `gorm:"not null" json:"first_name"`
	LastName    string    `gorm:"not null" json:"last_name"`
	Gender      string    `gorm:"not null" json:"gender"`
	Birthday    time.Time `gorm:"not null; type:date" json:"birthday"`
	Zipcode     string    `gorm:"not null" json:"zipcode"`
	Prefecture  string    `gorm:"not null" json:"prefecture"`
	City        string    `gorm:"not null" json:"city"`
	AddressLine string    `gorm:"not null" json:"address_line"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	User        User      `gorm:"foreignkey:user_id" json:"user"`
}
