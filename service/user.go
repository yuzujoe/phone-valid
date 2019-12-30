package service

import (
	"log"
	"phone-valid/models"
	"phone-valid/mysql"
)

func CreatePatient(phoneNumber string) string {

	db := mysql.DB

	user := models.User{PhoneNumber: phoneNumber}
	if err := db.FirstOrCreate(&user).Error; err != nil {
		log.Fatalln(err)
		return "phone"
	}

	return "true"
}

func RegisterCode(code string) error {
	db := mysql.DB

	var user models.User

	if err := db.Model(&user).Update("code", code).Error; err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
