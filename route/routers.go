package route

import (
	"phone-valid/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	rg := r.Group("/api/v1")
	v1 := rg.Group("v1")

	users := v1.Group("/users")
	{
		users.POST("/signup", controller.Signup)
		users.POST("/authentication", controller.Authentication)
		users.POST("/:userID/profile", controller.CreateProfile)
	}
}
