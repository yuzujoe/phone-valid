package controller

import (
	"log"
	"net/http"
	"phone-valid/util/request"

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
	if resp == nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Authentication user認証の関数s
func Authentication(c *gin.Context) {
	req := request.UserAuthenticationRequest{}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	resp, err := userAuthenticationImpl(c, &req)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
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
