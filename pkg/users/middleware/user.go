package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/userInterface"
)

func CheckRequiredFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requiredFields userInterface.Users
		if err := c.ShouldBindJSON(&requiredFields); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("missing required fields").Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
