package userInterface

type Users struct {
	ID            string `json:"id" bson:"_id"`
	NAME          string `json:"name" bson:"name" binding:"required"`
	EMAIL         string `json:"email"  bson:"email" binding:"required"`
	PASSWORD      string `json:"password"  bson:"password" binding:"required"`
	ISVERIFIED    bool   `json:"isVerified" bson:"isVerified"`
	CREATEDAT     int    `json:"created_at" bson:"created_at"`
	LASTUPDATEDAT int    `json:"last_updated_at" bson:"last_updated_at"`
}

type RegisterPayload struct {
	NAME     string `json:"name" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}
