package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
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

	fmt.Println(body)
}
