package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guopuke/longo/handler"
	"github.com/guopuke/longo/pkg/errno"
	"github.com/guopuke/longo/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
