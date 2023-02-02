package blogApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchUserBlogs(c *gin.Context) {
	id := c.Param("id")

	page := c.Request.URL.Query().Get("page")
	limit := c.Request.URL.Query().Get("limit")

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	// if !exist {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": constant.INVALID_REQUEST,
	// 	})
	// 	c.Abort()
	// 	return
	// }

	// serilisation
	user_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constant.PARSING_ERROR,
		})
		c.Abort()
		return
	}

	// checking blogs for user
	query := bson.M{
		"author": user_id,
		"status": constant.Published,
	}

	blogs := databaseConnector.FetchBlogs(query, page, limit)

	c.JSON(http.StatusOK, gin.H{
		"status": constant.SUCCESS,
		"blogs":  blogs,
	})
}
