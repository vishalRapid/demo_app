package blogApi

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"github.com/vishalrana9915/demo_app/pkg/algolia"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/redisConnector"
	"github.com/vishalrana9915/demo_app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to create a new blog
func CreateBlog(c *gin.Context) {

	// checking user required auth
	currentUserId, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userId not found in context"})
		return
	}

	// blog content
	var blogPayload blogInterface.Blog

	decoder := json.NewDecoder(c.Request.Body)

	// decode request body
	if err := decoder.Decode(&blogPayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		return
	}

	if blogPayload.TITLE == "" {
		fmt.Println("NO need to store as no title is there")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": constant.MISSING_TITLE})
		return
	}

	// creating check sum for blog
	blogPayload.CHECKSUM = utils.CreateCheckSum(blogPayload.CONTENT)
	blogPayload.WORDCOUNT = len(strings.Fields(blogPayload.CONTENT))
	blogPayload.READTIME = int(math.Round(float64(blogPayload.WORDCOUNT) / 250))
	blogPayload.ID = primitive.NewObjectID()
	// converting user id for author
	userID, err_author := primitive.ObjectIDFromHex(currentUserId.(string))
	if err_author != nil {
		panic(err_author)
	}

	objectId, _ := shortid.Generate()
	blogPayload.OBJECTID = objectId
	blogPayload.AUTHOR = userID
	blogPayload.CREATEDAT = time.Now()
	blogPayload.UPDATEDAT = time.Now()

	slug, err_slug := utils.MakeSlug(blogPayload.TITLE)

	if err_slug != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		return
	}

	blogPayload.SLUG = slug

	if blogPayload.READTIME == 0 {
		blogPayload.READTIME = 1
	}

	if blogPayload.STATUS == constant.Published {
		fmt.Println("we need to publish this blog")
		blogPayload.PUBLISHEDAT = time.Now()

		go redisConnector.MultipleSortedTags(blogPayload.TAGS)
	}

	// save blog in database
	go databaseConnector.CreateNewBlog(blogPayload)

	// save blog indexing
	go algolia.SaveBlogsData([]blogInterface.Blog{blogPayload})

	c.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{"data": blogPayload},
	)

}
