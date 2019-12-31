package service

import (
	"phone-valid/models"
	"phone-valid/mysql"
	"time"
)

func UserExist(phoneNumber string) *models.User {
	db := mysql.DB

	var user models.User

	auth := db.Where("phone_number = ?", phoneNumber).Select("user_id, phone_number").First(&user).RecordNotFound()

	if auth {
		return nil
	}

	return &user
}

func GetCodeInfo(phoneNumber string) *models.AuthenticationCode {
	db := mysql.DB

	var authentication models.AuthenticationCode

	authCode := db.Where("phone_number = ?", phoneNumber).Select("phone_number, code, expired").First(&authentication).RecordNotFound()

	if authCode {
		return nil
	}

	return &authentication
}

func CompareCode(code, reqCode string) bool {
	if code != reqCode {
		return false
	}
	return true
}

func CheckExpired(expired time.Time) (bool, error) {
	old := expired
	now := time.Now()

	if old.Before(now) {
		return false, nil
	}

	return true, nil
}
