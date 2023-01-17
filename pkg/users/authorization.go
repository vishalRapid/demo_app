package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
	"github.com/vishalrana9915/demo_app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// check if user exist
	userExist := databaseConnector.CheckUserExist(body.EMAIL)

	if userExist == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.USER_EXIST})
		c.Abort()
		return
	}

	// hashed password
	hashedPassword := utils.HashPassword(body.PASSWORD)

	// mapping Values
	body.PASSWORD = hashedPassword
	body.ISVERIFIED = false
	body.CREATEDAT = time.Now()
	body.LASTUPDATEDAT = time.Now()
	body.ID = primitive.NewObjectID()

	response := databaseConnector.CreateNewUser(body)

	c.JSON(http.StatusOK, gin.H{"message": constant.SUCCESS, "response": response})

}
