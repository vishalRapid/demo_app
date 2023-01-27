package databaseConnector

import (
	"context"
	"log"
	"strconv"

	"github.com/vishalrana9915/demo_app/pkg/blogs/blogInterface"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func FetchBlogs(filter interface{}, page string, limit string) []blogInterface.Blog {
	var blogs []blogInterface.Blog

	ctx := context.Background()

	// Get the page and limit query parameters
	page_no, _ := strconv.Atoi(page)
	limit_no, _ := strconv.Atoi(limit)

	if page_no < 0 {
		return blogs
	}

	if limit_no < 10 {
		return blogs
	}

	// Calculate the skip value
	skip := (page_no - 1) * limit_no

	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit_no))

	cur, err := Adapter.db.Collection(constant.BLOGCOLLECTION).Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result blogInterface.Blog
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		result.CONTENT = result.CONTENT[:100]

		blogs = append(blogs, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return blogs

}
