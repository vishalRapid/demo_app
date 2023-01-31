package blogApi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogsUtils"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"go.mongodb.org/mongo-driver/bson"
)

// function to fetch bogs based on user input
// recommendation or based out of a tag
func FetchBlogs(c *gin.Context) {

	userId, exist := c.Get("userId")

	querySetting := c.Request.URL.Query().Get("search")
	page := c.Request.URL.Query().Get("page")
	limit := c.Request.URL.Query().Get("limit")

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	var searchCriteria string

	if !exist && querySetting == "" {
		fmt.Print("No user information available")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constant.INVALID_REQUEST,
		})
		c.Abort()
		return
	}

	if querySetting == "" {
		searchCriteria = constant.Recommendation
	} else {
		searchCriteria = querySetting
	}

	if searchCriteria == constant.Recommendation {
		blogs := blogsUtils.GetRecommendedBlogs(userId.(string), page, limit)
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": constant.SUCCESS,
				"blogs":   blogs,
			},
		)
		return
	}

	query := bson.M{
		"tags":   searchCriteria,
		"status": constant.Published,
	}

	blogs := databaseConnector.FetchBlogs(query, page, limit)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": constant.SUCCESS,
			"blogs":   blogs,
		},
	)

}
