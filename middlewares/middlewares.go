
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"web/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
func JwtAuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValidAdmin(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
