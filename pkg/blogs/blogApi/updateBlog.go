package blogApi

import (
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/redisConnector"
	"github.com/vishalrana9915/demo_app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to update details of a blog
func UpdateBlog(c *gin.Context) {

	// checking user required auth
	currentUserId, exist := c.Get("userId")

	slug := c.Param("slug")

	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userId not found in context"})
		return
	}

	// serializing request body
	var blogPayload blogInterface.UpdateBlog

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&blogPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	userID, err_author := primitive.ObjectIDFromHex(currentUserId.(string))
	if err_author != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	// fetch blog from database
	query := bson.M{
		"slug":   slug,
		"author": userID,
	}

	currentBlog := databaseConnector.FetchBlogQuery(query)

	if currentBlog.SLUG == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.BLOG_NOT_FOUND})
		c.Abort()
		return
	}

	// mapping user request payload with database updater
	updatedBlog := blogInterface.DatabaseBlog{
		SLUG:         currentBlog.SLUG,
		CHECKSUM:     currentBlog.CHECKSUM,
		WORDCOUNT:    currentBlog.WORDCOUNT,
		READTIME:     currentBlog.READTIME,
		TITLE:        blogPayload.TITLE,
		CONTENT:      blogPayload.CONTENT,
		BANNER:       blogPayload.BANNER,
		TAGS:         blogPayload.TAGS,
		STATUS:       blogPayload.STATUS,
		FEATUREIMAGE: blogPayload.FEATUREIMAGE,
	}

	// checking if title has changed or not
	if currentBlog.TITLE != updatedBlog.TITLE || currentBlog.SLUG == "" {

		// need to update title and slug
		updated_slug, err_slug := utils.MakeSlug(updatedBlog.TITLE)

		if err_slug != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
			c.Abort()
			return
		}
		updatedBlog.SLUG = updated_slug
	}

	// checking the content of the blog, if it is been updated or not
	new_checksum := utils.CreateCheckSum(updatedBlog.CONTENT)

	if new_checksum != currentBlog.CHECKSUM {
		// content has changed, need to update word count and read time
		updatedBlog.WORDCOUNT = len(strings.Fields(updatedBlog.CONTENT))
		updatedBlog.READTIME = int(math.Round(float64(updatedBlog.WORDCOUNT) / 250))
		updatedBlog.CHECKSUM = new_checksum
		if updatedBlog.READTIME == 0 {
			updatedBlog.READTIME = 1
		}
	}

	// checking tags
	if len(updatedBlog.TAGS) > 5 {
		updatedBlog.TAGS = updatedBlog.TAGS[:5]
	}

	// checking if status has changed for this or not
	if currentBlog.STATUS != updatedBlog.STATUS {
		// status has changed
		if updatedBlog.STATUS == constant.Published {
			// current status is changed to published, need to update tags
			updatedBlog.PUBLISHEDAT = time.Now()

			go redisConnector.MultipleSortedTags(updatedBlog.TAGS)
		}
	}

	// update content of the blog
	updates := bson.M{
		"$set": updatedBlog,
	}

	_, err_updateBlog := databaseConnector.UpdateBlog(query, updates)

	if err_updateBlog != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constant.INVALID_REQUEST,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": constant.SUCCESS,
	})
}
