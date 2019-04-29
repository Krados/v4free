package routes

import (
	"app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	// -------------------------------------------------
	// global middlewares
	// -------------------------------------------------
	// reload .env everytime
	app.Use(middlewares.ReloadEnv())

	// check whether server is in maintenance mode or not
	app.Use(middlewares.CheckForMaintenance())

	// log every request
	app.Use(middlewares.LogRequest())

	// -------------------------------------------------
	// custom routes
	// -------------------------------------------------
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "Posts",
		})
	})
}
