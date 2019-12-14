package route

import (
	"github.com/gin-gonic/gin"
	"phone-valid/controller"
)

func Route(r *gin.Engine) {

	rg := r.Group("/api")
	v1 := rg.Group("v1")
	{
		v1.POST("/phone", controller.Signup)
	}
}
