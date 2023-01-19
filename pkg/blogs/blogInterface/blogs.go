package blogInterface

import "time"

type Blog struct {
	ID           string    `json:"id" bson:"id"`
	TITLE        string    `json:"title" bson:"title"`
	CONTENT      string    `json:"content" bson:"content"`
	BANNER       string    `json:"banner" bson:"banner"`
	AUTHOR       string    `json:"author" bson:"author"`
	CHECKSUM     string    `json:"checksum" bson:"checksum"`
	WORDCOUNT    int       `json:"wordcount" bson:"wordcount"`
	TAGS         []string  `json:"tags" bson:"tags"`
	CREATEDAT    time.Time `json:"created_at" bson:"created_at"`
	UPDATEDAT    time.Time `json:"updated_at" bson:"updated_at"`
	PUBLISHEDAT  time.Time `json:"published_at" bson:"published_at"`
	STATUS       int       `json:"status" bson:"status"`
	FEATUREIMAGE string    `json:"featureImage" bson:"featureImage"`
}
