package commonMiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AUth guard to check required auth token for validation
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// Do something with the token

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": constant.UNAUTH_REQUEST,
			})
			c.Abort()
			return
		}

		// validate token
		userId, err := utils.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": constant.UNAUTH_REQUEST,
			})
			c.Abort()
			return
		}

		// fetch userId details from database

		objectID, err_objectId := primitive.ObjectIDFromHex(userId)
		if err_objectId != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": constant.UNAUTH_REQUEST,
			})
			c.Abort()
			return
		}
		// check if user exist
		query := bson.D{
			{
				"_id", objectID,
			},
		}
		_, errr := databaseConnector.FindUser(query)
		if errr != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": errr})
			c.Abort()
			return
		}

		c.Set("userId", userId)
	}
}
