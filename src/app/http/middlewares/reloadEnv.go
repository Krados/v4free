package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ReloadEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Overload("../.env")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"message": "env 檔不存在",
			})
		}
	}
}
