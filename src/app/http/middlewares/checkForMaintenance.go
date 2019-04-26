package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckForMaintenance() gin.HandlerFunc {
	return func(c *gin.Context) {
		envVariable := os.Getenv("IS_ONLINE")
		if envVariable != "true" {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"message": "server is in maintenance mode",
			})
		}
	}
}
