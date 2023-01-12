package users

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/userInterface"
)

// a function handler to take care of user onboarding
func RegisterUser(c *gin.Context) {
	var payload userInterface.Users
	json.Unmarshal(c.Request.Body, &payload)
	fmt.Println(payload)

}
