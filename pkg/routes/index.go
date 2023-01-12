package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Setting up router for the app
func SetupRouter(router *gin.Engine) {

	proxiers := os.Getenv("TRUSTED_PROXIES")

	router.SetTrustedProxies([]string{proxiers})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

}
