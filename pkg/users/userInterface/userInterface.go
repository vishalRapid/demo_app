package userInterface

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	NAME          string             `json:"name" bson:"name" binding:"required"`
	EMAIL         string             `json:"email"  bson:"email" binding:"required"`
	PASSWORD      string             `json:"password"  bson:"password" binding:"required"`
	ISVERIFIED    bool               `json:"isVerified" bson:"isVerified"`
	TAGS          []string           `json:"tags" bson:"tags"`
	CREATEDAT     time.Time          `json:"created_at" bson:"created_at"`
	LASTUPDATEDAT time.Time          `json:"last_updated_at" bson:"last_updated_at"`
}

type RegisterPayload struct {
	NAME     string `json:"name" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

type LoginPayload struct {
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

type UserProfile struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	NAME          string             `json:"name" bson:"name"`
	EMAIL         string             `json:"email"  bson:"email"`
	ISVERIFIED    bool               `json:"isVerified" bson:"isVerified"`
	CREATEDAT     time.Time          `json:"created_at" bson:"created_at"`
	LASTUPDATEDAT time.Time          `json:"last_updated_at" bson:"last_updated_at"`
}

type UserTags struct {
	TAGS []string `json:"tags" bson:"tags"`
}
