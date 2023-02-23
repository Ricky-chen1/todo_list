package middleware

import (
	"time"
	"westonline/utilities/tokenfunc"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		var msg string
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			msg = "token为空"
		} else {
			claim, err := tokenfunc.ParseToken(token, "golang")
			if err != nil {
				code = 403
				msg = "用户无权限"
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401
				msg = "token已过期"
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
