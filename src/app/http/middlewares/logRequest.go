package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("log your request in somewhere")
	}
}
