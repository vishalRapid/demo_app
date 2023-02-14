package blogApi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get api for blog to be fetched from slug
func FetchBlog(c *gin.Context) {

	userId, exist := c.Get("userId")

	// checking if slug is present or not
	slug := c.Param("slug")

	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constant.MISSING_SLUG,
		})
	}
	query := bson.M{
		"slug":   slug,
		"status": constant.Published,
	}

	// database query to fetch blog details
	blog := databaseConnector.FetchBlogQuery(query)

	// setting up user history
	if exist {

		userObjectId, err := primitive.ObjectIDFromHex(userId.(string))

		if err != nil {
			fmt.Println("Error saving user history", err)
		} else {
			blogHistory := blogInterface.BlogHistory{
				ID:        primitive.NewObjectID(),
				AUTHOR:    blog.AUTHOR,
				BLOGID:    blog.ID,
				TAGS:      blog.TAGS,
				USERID:    userObjectId,
				CREATEDAT: time.Now(),
			}

			go databaseConnector.CreateHistory(blogHistory)
		}

	}

	// increment read count for the blog

	// setting up cache for current blog
	c.Header("Cache-Control", "public, max-age=3600")

	c.JSON(
		http.StatusOK, gin.H{
			"message": constant.SUCCESS,
			"data":    blog,
		},
	)

}
