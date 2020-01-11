package controller

import (
	"log"
	"net/http"
	"phone-valid/util/jwt"
	"phone-valid/util/request"
	"strconv"
	"time"

	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Signup(c *gin.Context) {
	req := request.UserSignupRequest{}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	resp := userSignupImpl(c, &req)
	if resp != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
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

func CreateProfile(c *gin.Context) {
	var req request.CreateProfileRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	c.JSON(http.StatusOK, "ok")
}
