package users

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
	"github.com/vishalrana9915/demo_app/pkg/utils"
)

// a function handler to take care of user onboarding
func RegisterUser(c *gin.Context) {
	var body userInterface.Users
	decode := json.NewDecoder(c.Request.Body)
	if err := decode.Decode(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	// hashed password
	hashedPassword := utils.HashPassword(body.PASSWORD)

	body.PASSWORD = hashedPassword

}
