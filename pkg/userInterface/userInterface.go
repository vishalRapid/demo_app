package userInterface

type Users struct {
	ID            string `json:"id" `
	NAME          string `json:"name" binding:"required"`
	EMAIL         string `json:"email" binding:"required"`
	ISVERIFIED    bool   `json:"isverified"`
	CREATEDAT     int    `json:"created_at"`
	LASTUPDATEDAT int    `json:"last_updated_at"`
}
