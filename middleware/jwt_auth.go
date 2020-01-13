package middleware

import (
	"log"
	"net/http"
	"phone-valid/models"
	"phone-valid/mysql"
	"phone-valid/util/jwt"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserID      int64
	PhoneNumber string
}

func JwtAuth(c *gin.Context) {
	if c.Request.URL.Path == "/api/v1/users/signup" || c.Request.URL.Path == "/api/v1/users/authentication" {
		c.Next()
	} else {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			log.Println("empty token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		} else if payload, error := jwt.TokenParse(token); error != nil {
			log.Println(error)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		} else {
			db := mysql.DB
			var user models.User
			if userInfo := db.Select("user_id, phone_number").Where("user_id = ?", payload["UserID"]).First(&user).RecordNotFound(); userInfo {
				log.Println(userInfo)
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
			}

			authInfo := UserInfo{
				user.UserID,
				user.PhoneNumber,
			}

			c.Set("userID", payload["UserID"])
			c.Set("authInfo", authInfo)

			c.Next()
		}
	}
}
