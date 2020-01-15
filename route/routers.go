package route

import (
	"github.com/gin-gonic/gin"
	"phone-valid/internal/user"
)

func Route(r *gin.Engine) {

	rg := r.Group("/api/v1")

	users := rg.Group("/users")
	{
		users.POST("/signup", user.Signup)
		users.POST("/authentication", user.Authentication)
		users.POST("/profile/:userID", user.CreateProfile)
	}
}
