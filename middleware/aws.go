package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws/session"
)

func Aws() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		if sess == nil {
			c.Abort()
			return
		}

		c.Set("AwsSession", sess)
	}
}
