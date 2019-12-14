package controller

import (
	"log"
	"net/http"
	"phone-valid/models"
	"phone-valid/mysql"
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

	err := phoneChk(req.PhoneNumber)
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "他に電話番号が登録されていました",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// PhoneChk 電話番号が他に存在しないかチェックする関数
func phoneChk(phoneNumber string) bool {
	db := mysql.DB
	var user models.User
	phoneChk := db.Select("phone_number").Where("phone_number = ?", phoneNumber).First(&user).RecordNotFound()

	return phoneChk
}
