package service

import (
	"log"
	"phone-valid/models"
	"phone-valid/mysql"
	"time"
)

func CreatePatient(phoneNumber string) bool {

	db := mysql.DB

	var user models.User
	if err := db.FirstOrCreate(&user, models.User{PhoneNumber: phoneNumber}).Error; err != nil {
		log.Fatalln(err)
		return false
	}

	return true
}

func RegisterCode(phoneNumber, code string) error {
	db := mysql.DB

	var authCode models.AuthenticationCode

	expired := time.Now().Add(15 * time.Minute)

	if err := db.Where(models.AuthenticationCode{PhoneNumber: phoneNumber}).Assign(models.AuthenticationCode{Code: code, Expired: expired, UpdatedAt: time.Now()}).FirstOrCreate(&authCode).Error; err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
