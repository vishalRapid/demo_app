package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
	"gopkg.in/go-playground/validator.v9"
)

// function to validate if email is valid or not
func validateEmail(email string) bool {
	validator := validator.New()
	err := validator.Var(email, "required,email")
	if err != nil {
		return false
	} else {
		return true
	}
}

// request middleware to check if all required fields are present or not for this user
func CheckRequiredFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requiredFields userInterface.RegisterPayload
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constant.MISSING_REQUIRED_FIELDS,
			})
			c.Abort()
			return
		}
		if err = json.Unmarshal(body, &requiredFields); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constant.MISSING_REQUIRED_FIELDS,
			})
			c.Abort()
			return
		}

		// validating requrired email value
		if !validateEmail(requiredFields.EMAIL) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constant.INVALID_EMAIL,
			})
			c.Abort()
			return
		}

		// checking name values for missing
		if requiredFields.NAME == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constant.MISSING_NAME,
			})
			c.Abort()
			return
		}

		// checking password values for missing
		if requiredFields.PASSWORD == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": constant.MISSING_PASSWORD,
			})
			c.Abort()
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()
	}
}
