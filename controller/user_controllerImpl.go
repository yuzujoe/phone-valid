package controller

import (
	"log"
	"phone-valid/models"
	"phone-valid/mysql"
	"time"
)

func createPatient(phoneNumber string) bool {

	db := mysql.DB

	var user models.User
	if err := db.FirstOrCreate(&user, models.User{PhoneNumber: phoneNumber}).Error; err != nil {
		log.Fatalln(err)
		return false
	}

	return true
}

func registerCode(phoneNumber, code string) error {
	db := mysql.DB

	var authCode models.AuthenticationCode

	expired := time.Now().Add(15 * time.Minute)

	if err := db.Where(models.AuthenticationCode{PhoneNumber: phoneNumber}).Assign(models.AuthenticationCode{Code: code, Expired: expired, UpdatedAt: time.Now()}).FirstOrCreate(&authCode).Error; err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func userExist(phoneNumber string) *models.User {
	db := mysql.DB

	var user models.User

	auth := db.Where("phone_number = ?", phoneNumber).Select("user_id, phone_number").First(&user).RecordNotFound()

	if auth {
		return nil
	}

	return &user
}

func getCodeInfo(phoneNumber string) *models.AuthenticationCode {
	db := mysql.DB

	var authentication models.AuthenticationCode

	authCode := db.Where("phone_number = ?", phoneNumber).Select("phone_number, code, expired").First(&authentication).RecordNotFound()

	if authCode {
		return nil
	}

	return &authentication
}

func compareCode(code, reqCode string) bool {
	if code != reqCode {
		return false
	}
	return true
}

func checkExpired(expired time.Time) (bool, error) {
	old := expired
	now := time.Now()

	if old.Before(now) {
		return false, nil
	}

	return true, nil
}
