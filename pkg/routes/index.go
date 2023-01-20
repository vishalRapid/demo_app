package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogApi"
	"github.com/vishalrana9915/demo_app/pkg/mediaUpload"
	"github.com/vishalrana9915/demo_app/pkg/users"
	"github.com/vishalrana9915/demo_app/pkg/users/middleware"
	"github.com/vishalrana9915/demo_app/pkg/utils/commonMiddleware"
)

// Setting up router for the app
func SetupRouter(router *gin.Engine) {

	proxiers := os.Getenv("TRUSTED_PROXIES")

	router.SetTrustedProxies([]string{proxiers})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	//////////////////////////////// User routes //////////////////////////////
	router.POST("/register", middleware.CheckRequiredFields(), users.RegisterUser)

	router.GET("/login", users.AuthenticateUser)

	router.GET("/profile/me", users.FetchProfile)

	// blogs routes

	router.POST("/blogs/create", commonMiddleware.AuthGuard(), blogApi.CreateBlog)

	// upload media

	router.PUT("/media/upload", commonMiddleware.AuthGuard(), mediaUpload.UploadMedia)

}
