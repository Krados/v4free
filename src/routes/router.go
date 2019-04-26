package routes

import (
	"app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// -------------------------------------------------
	// global middlewares
	// -------------------------------------------------
	// reload .env everytime
	router.Use(middlewares.ReloadEnv())

	// check whether server is in maintenance mode or not
	router.Use(middlewares.CheckForMaintenance())

	// log every request
	router.Use(middlewares.LogRequest())

	// -------------------------------------------------
	// custom routes
	// -------------------------------------------------
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "Posts",
		})
	})
}
