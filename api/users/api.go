package users

import (
	"phone-valid/controller"

	"github.com/gin-gonic/gin"
)
// Api users api
func Api(r *gin.RouterGroup) {
	g := r.Group("/users")
	{
		g.POST("/signup", controller.Signup)
		g.POST("/authentication", controller.Authentication)
	}
}
