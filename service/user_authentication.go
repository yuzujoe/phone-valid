package service

import (
	"phone-valid/models"
	"phone-valid/mysql"
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
