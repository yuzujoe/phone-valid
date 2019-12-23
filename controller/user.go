package controller

import (
	"crypto/rand"
	"io"
	"log"
	"net/http"
	"phone-valid/models"
	"phone-valid/mysql"
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

	code := generateAuthCode(codeLength)

	if err := createPatient(req.PhoneNumber); err != "true" {
		log.Fatalln(err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "error",
		})
		return
	}

	if err := registerCode(code); err != nil {
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

	code := generateAuthCode(codeLength)

	if err := createPatient(req.PhoneNumber, code); err != "true" {
		log.Fatalln(err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "error",
		})
		return
	}

	log.Println("ok")

	sms.PushSms(req.PhoneNumber, code)

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// generateAuthCode 認証コード作成のロジック
func generateAuthCode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return err.Error()
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func createPatient(phoneNumber string) string {

	db := mysql.DB

	user := models.User{PhoneNumber: phoneNumber}
	if err := db.FirstOrCreate(&user).Error; err != nil {
		log.Fatalln(err)
		return "phone"
	}

	return "true"
}

func registerCode(code string) error {
	db := mysql.DB

	var user models.User

	if err := db.Model(&user).Update("code", code).Error; err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

// generateAuthCode 認証コード作成のロジック
func generateAuthCode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return err.Error()
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func createPatient(phoneNumber, code string) string {

	db := mysql.DB
	tx := db.Begin()

	user := models.User{PhoneNumber: phoneNumber, Code: code}
	if err := tx.Create(&user).Error; err != nil {
		log.Fatalln(err)
		return "db"
	}

	tx.Commit()

	return "true"
}
