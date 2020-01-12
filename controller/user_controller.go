package controller

import (
	"log"
	"net/http"
	"phone-valid/util/request"
	"phone-valid/util/response"

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
	req := request.CreateProfileRequest{}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    400,
			Message: "Bad Request",
		})
		return
	}

	resp, err := userProfileCreateImpl(c, &req)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}
