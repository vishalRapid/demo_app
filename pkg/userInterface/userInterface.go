package userInterface

type Users struct {
	ID            string `json:"id"`
	NAME          string `json:"name"`
	EMAIL         string `json:"email"`
	ISVERIFIED    bool   `json:"isverified"`
	CREATEDAT     int    `json:"created_at"`
	LASTUPDATEDAT int    `json:"last_updated_at"`
}
