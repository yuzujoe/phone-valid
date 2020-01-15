package controller

import (
	"log"
	"net/http"
	"phone-valid/models"
	"phone-valid/mysql"
	"phone-valid/util/request"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func insertProfile(c *gin.Context, request *request.CreateProfileRequest) error {
	db := mysql.DB

	userIDStr, _ := c.Get("userID")
	userID, err := strconv.ParseInt(userIDStr.(string), 10, 64)
	if err != nil {
		log.Println(err)
		return err
	}

	userProfile := models.UserProfile{UserID: userID, Email: request.Email, FirstName: request.FirstName, LastName: request.LastName, Gender: request.Gender, Birthday: request.Birthday, Zipcode: request.Zipcode, Prefecture: request.Prefecture, City: request.City, AddressLine: request.AddressLine, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if err := db.Create(&userProfile); err != nil {
		log.Println(err)
		return err.Error
	}

	return nil
}

func createUser(phoneNumber string) error {

	db := mysql.DB

	var user models.User
	if err := db.FirstOrCreate(&user, models.User{PhoneNumber: phoneNumber}).Error; err != nil {
		log.Fatalln(err)
		return nil
	}

	return nil
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

func userExist(phoneNumber string) (models.User, error) {
	db := mysql.DB

	var user models.User

	if err := db.Where("phone_number = ?", phoneNumber).Select("user_id, phone_number").First(&user).Error; err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func getCodeInfo(phoneNumber string) (*models.AuthenticationCode, error) {
	db := mysql.DB

	var authentication models.AuthenticationCode

	authCode := db.Where("phone_number = ?", phoneNumber).Select("phone_number, code, expired").First(&authentication).RecordNotFound()

	if authCode {
		return nil, db.Error
	}

	return &authentication, nil
}

func compareCode(code, reqCode string) error {
	if code != reqCode {
		return http.ErrAbortHandler
	}
	return nil
}

func checkExpired(expired time.Time) error {
	old := expired
	now := time.Now()

	if old.Before(now) {
		return http.ErrAbortHandler
	}

	return nil
}
