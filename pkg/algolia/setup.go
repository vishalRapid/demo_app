package algolia

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
)

var Adapter = AlgoliaAdapter{}

type AlgoliaAdapter struct {
	client    *search.Client
	userIndex *search.Index
	blogIndex *search.Index
}

// setup algolia for searching
func (a *AlgoliaAdapter) SetupAlgolia() {

	// setup search
	var appId string = os.Getenv("ALGOLIA_APP_ID")
	var apiKey string = os.Getenv("ALGOLIA_ADMIN_API_KEY")

	client := search.NewClient(appId, apiKey)

	Adapter.client = client

	// setting up indexes

	blogIndex := os.Getenv("ALOGIA_BLOG_INDEX")

	index := Adapter.client.InitIndex(blogIndex)
	Adapter.blogIndex = index

	// setting up user index

	userIndex := os.Getenv("ALGOLIA_USER_INDEX")

	user_index := Adapter.client.InitIndex(userIndex)
	Adapter.userIndex = user_index
}

// save data to the index
func SaveBlogsData(blog []blogInterface.Blog) search.GroupBatchRes {

	res, err := Adapter.blogIndex.SaveObjects(blog, opt.AutoGenerateObjectIDIfNotExist(true))

	if err != nil {
		panic(err)
	}

	return res
}

// batch update user data to algolia for indexing
func saveUsersData(user userInterface.UserProfile) search.GroupBatchRes {

	res, err := Adapter.userIndex.SaveObjects(user, opt.AutoGenerateObjectIDIfNotExist(true))

	if err != nil {
		panic(err)
	}

	return res
}
