package blogsUtils

import (
	"encoding/json"
	"fmt"

	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/redisConnector"
	"go.mongodb.org/mongo-driver/bson"
)

/**
* Function to fetch recommended blogs for current user
 */
func GetRecommendedBlogs(userId string, page string, limit string) []blogInterface.Blog {

	// fetch user information
	userTags, err := redisConnector.GetValue(fmt.Sprintf("user_tags_%s", userId))

	if err != nil {
		panic(err)

	}

	var tags []string
	if err_tags := json.Unmarshal([]byte(userTags), &tags); err_tags != nil {
		panic(err_tags)
	}

	// need to fetch latest blogs that matched to this user tags preferrence
	query := bson.M{
		"tags": bson.M{
			"$in": tags,
		},
		"status": constant.Published,
	}

	blogs := databaseConnector.FetchBlogs(query, page, limit)

	return blogs
}
