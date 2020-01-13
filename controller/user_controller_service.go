package controller

import (
	"log"
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
