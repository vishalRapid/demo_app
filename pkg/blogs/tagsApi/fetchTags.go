package tagsApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/redisConnector"
)

// api for user to fetch available tags in the database to choose
func FetchTags(c *gin.Context) {

	tags := redisConnector.GetAllSet("tags")

	// if len(tags) > 20 {
	// 	tags = tags[:20]
	// }

	c.JSON(http.StatusOK, gin.H{
		"tags": tags,
	})
}
