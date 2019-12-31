package route

import (
	"phone-valid/api/users"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	rg := r.Group("/api")
	v1 := rg.Group("v1")
	{
		users.Api(v1)
	}
}
