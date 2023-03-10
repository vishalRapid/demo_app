package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogApi"
	"github.com/vishalrana9915/demo_app/pkg/blogs/tagsApi"
	"github.com/vishalrana9915/demo_app/pkg/mediaUpload"
	"github.com/vishalrana9915/demo_app/pkg/users"
	"github.com/vishalrana9915/demo_app/pkg/users/middleware"
	"github.com/vishalrana9915/demo_app/pkg/utils"
	"github.com/vishalrana9915/demo_app/pkg/utils/commonMiddleware"
)

// Setting up router for the app
func SetupRouter(router *gin.Engine) {

	proxiers := os.Getenv("TRUSTED_PROXIES")

	router.SetTrustedProxies([]string{proxiers})

	router.Use(utils.AssignRequestID)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	//////////////////////////////// User routes //////////////////////////////
	router.POST("/register", middleware.CheckRequiredFields(), users.RegisterUser)

	router.GET("/login", users.AuthenticateUser)

	router.GET("/profile/me", users.FetchProfile)

	router.PUT("/profile/tags", commonMiddleware.AuthGuard(), users.UpdateTags)

	// updating profile for user
	router.PUT("/profile/me", commonMiddleware.AuthGuard(), users.UpdateProfile)

	// blogs routes

	router.POST("/blogs/create", commonMiddleware.AuthGuard(), blogApi.CreateBlog)

	router.GET("/blogs/:slug", blogApi.FetchBlog)

	router.PUT("/blogs/:id", commonMiddleware.AuthGuard(), blogApi.UpdateBlog)

	router.GET("/tags", tagsApi.FetchTags)

	router.GET("/blogs/timeline", commonMiddleware.OptionalAuthGuard(), blogApi.FetchBlogs)

	// fetch user blogs
	router.GET("/user/blogs/:id", blogApi.FetchUserBlogs)

	// upload media

	router.PUT("/media/upload", commonMiddleware.AuthGuard(), mediaUpload.UploadMedia)

}
