package databaseConnector

import (
	"context"

	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"go.mongodb.org/mongo-driver/mongo"
)

// save blogs record into database
func CreateNewBlog(data interface{}) interface{} {

	result, err := Adapter.db.Collection(constant.BLOGCOLLECTION).InsertOne(context.TODO(), data)

	if err != nil {
		panic(error.Error(err))
	}

	return result
}

// fetch blog details from query
func FetchBlogQuery(filter interface{}) blogInterface.Blog {

	var blog blogInterface.Blog
	err := Adapter.db.Collection(constant.BLOGCOLLECTION).FindOne(context.TODO(), filter).Decode(&blog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return blog
		}
		panic(error.Error(err))
	}

	return blog

}
