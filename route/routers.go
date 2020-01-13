package route

import (
	"phone-valid/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	rg := r.Group("/api/v1")

	users := rg.Group("/users")
	{
		users.POST("/signup", controller.Signup)
		users.POST("/authentication", controller.Authentication)
		users.POST("/profile/:userID", controller.CreateProfile)
	}
}
