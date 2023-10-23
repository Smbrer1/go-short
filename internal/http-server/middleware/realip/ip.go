package realip

import "github.com/gin-gonic/gin"

func RealIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ip := c.ClientIP(); ip != "" {
			c.Request.RemoteAddr = ip
		}
		c.Next()
	}
}
