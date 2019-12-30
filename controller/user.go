package controller

import (
	"log"
	"net/http"
	"phone-valid/service"
	"phone-valid/util/auth"
	"phone-valid/util/sms"
	"regexp"

	"github.com/gin-gonic/gin"
)

const codeLength = 6

func Signup(c *gin.Context) {

	type request struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}

	var req request

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	// 電話番号の正規表現
	policy := "^\\d{2,4}-?\\d{2,4}-?\\d{3,4}$"
	re := regexp.MustCompile(policy)
	reg := re.MatchString(req.PhoneNumber)
	if !reg {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	code := auth.GenerateAuthCode(codeLength)

	if err := service.CreatePatient(req.PhoneNumber); err != "true" {
		log.Fatalln(err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "error",
		})
		return
	}

	if err := service.RegisterCode(req.PhoneNumber, code); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "error",
		})
		return
	}

	if err := sms.PushSms(req.PhoneNumber, code); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "通信に失敗しました",
		})
		return
	}

	log.Println("ok")

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
