package controller

import (
	"errors"
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
		err = errors.New("please enter phone number")
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	resp, err := userSignupImpl(c, &req)
	if err != nil {
		log.Println(err)
		if errors.Is(err, response.UserSignup400Reponse) {
			c.JSON(400, &response.Response{
				Code:    400,
				Message: err.Error(),
			})
			return
		}
		c.JSON(500, &response.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Authentication user認証の関数s
func Authentication(c *gin.Context) {
	req := request.UserAuthenticationRequest{}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		err = response.UserAuthentication400Response
		log.Println(err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	resp, err := userAuthenticationImpl(c, &req)
	if err != nil {
		log.Println(err)
		if errors.Is(err, response.UserAuthentication400Response) {
			c.JSON(http.StatusBadRequest, &response.Response{
				Code:    400,
				Message: err.Error(),
			})
			return
		} else if errors.Is(err, response.UserAuthentication403Response) {
			c.JSON(http.StatusForbidden, &response.Response{
				Code:    403,
				Message: err.Error(),
			})
			return
		} else if errors.Is(err, response.UserAuthentication404Response) {
			c.JSON(http.StatusNotFound, &response.Response{
				Code:    404,
				Message: err.Error(),
			})
			return
		}
		c.JSON(500, &response.Response{
			Code:    500,
			Message: err.Error(),
		})
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
			Message: err.Error(),
		})
		return
	}

	resp, err := userProfileCreateImpl(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(500, &response.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
