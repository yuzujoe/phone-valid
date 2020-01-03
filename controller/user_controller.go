package controller

import (
	"log"
	"net/http"
	"phone-valid/util/auth"
	"phone-valid/util/jwt"
	"phone-valid/util/request"
	"phone-valid/util/sms"
	"regexp"
	"strconv"
	"time"

	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const codeLength = 6

func Signup(c *gin.Context) {
	var req request.UserSignupRequest

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

	if err := createPatient(req.PhoneNumber); err != true {
		log.Fatalln(err)
		c.JSON(http.StatusConflict, gin.H{
			"message": "error",
		})
		return
	}

	if err := registerCode(req.PhoneNumber, code); err != nil {
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

// Authentication user認証の関数s
func Authentication(c *gin.Context) {
	var req request.UserAuthenticationRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
		})
		return
	}

	user := userExist(req.PhoneNumber)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
	}

	authCode := getCodeInfo(req.PhoneNumber)
	if authCode == nil {
		return
	}

	compare := compareCode(authCode.Code, req.Code)
	if !compare {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "入力された認証コードが間違っています、再度正しい認証コードを取得してください",
		})
		return
	}

	expired, err := checkExpired(authCode.Expired)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	if !expired {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "有効期限が切れています、再度認証コードを取得してから入力してください",
		})
		return
	}

	userID := strconv.FormatInt(int64(user.UserID), 10)

	token, _ := jwt.TokenGenerate(jwt_go.MapClaims{
		"UserID": userID,
		"expire": time.Now().Add(time.Hour * 2).Unix(),
	})
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
