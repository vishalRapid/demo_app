package blogInterface

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	OBJECTID     string             `json:"objectID" bson:"objectID"`
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	SLUG         string             `json:"slug" bson:"slug"`
	TITLE        string             `json:"title" bson:"title"`
	CONTENT      string             `json:"content" bson:"content"`
	BANNER       string             `json:"banner" bson:"banner"`
	AUTHOR       primitive.ObjectID `json:"author" bson:"author"`
	CHECKSUM     string             `json:"checksum" bson:"checksum"`
	WORDCOUNT    int                `json:"wordcount" bson:"wordcount"`
	READTIME     int                `json:"readTime" bson:"readTime"`
	TAGS         []string           `json:"tags" bson:"tags"`
	CREATEDAT    time.Time          `json:"created_at" bson:"created_at"`
	UPDATEDAT    time.Time          `json:"updated_at" bson:"updated_at"`
	PUBLISHEDAT  time.Time          `json:"published_at" bson:"published_at"`
	STATUS       int                `json:"status" bson:"status"`
	FEATUREIMAGE string             `json:"featureImage" bson:"featureImage"`
}

type UpdateBlog struct {
	TITLE        string   `json:"title" bson:"title"`
	CONTENT      string   `json:"content" bson:"content"`
	BANNER       string   `json:"banner" bson:"banner"`
	TAGS         []string `json:"tags" bson:"tags"`
	STATUS       int      `json:"status" bson:"status"`
	FEATUREIMAGE string   `json:"featureImage" bson:"featureImage"`
}

type DatabaseBlog struct {
	SLUG         string    `json:"slug" bson:"slug"`
	TITLE        string    `json:"title" bson:"title"`
	CONTENT      string    `json:"content" bson:"content"`
	BANNER       string    `json:"banner" bson:"banner"`
	CHECKSUM     string    `json:"checksum" bson:"checksum"`
	WORDCOUNT    int       `json:"wordcount" bson:"wordcount"`
	READTIME     int       `json:"readTime" bson:"readTime"`
	TAGS         []string  `json:"tags" bson:"tags"`
	PUBLISHEDAT  time.Time `json:"published_at" bson:"published_at"`
	STATUS       int       `json:"status" bson:"status"`
	FEATUREIMAGE string    `json:"featureImage" bson:"featureImage"`
}
