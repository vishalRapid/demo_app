package blogApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"go.mongodb.org/mongo-driver/bson"
)

// get api for blog to be fetched from slug
func FetchBlog(c *gin.Context) {

	// checking if slug is present or not
	slug := c.Param("slug")

	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constant.MISSING_SLUG,
		})
	}
	query := bson.D{
		{
			"slug", slug,
		},
	}

	// database query to fetch blog details
	blog := databaseConnector.FetchBlogQuery(query)

	// setting up cache for current blog
	c.Header("Cache-Control", "public, max-age=3600")

	c.JSON(
		http.StatusOK, gin.H{
			"message": constant.SUCCESS,
			"data":    blog,
		},
	)

}
