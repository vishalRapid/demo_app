package responseHandler

import "github.com/gin-gonic/gin"

/*
*
Function to handle out the response
*/
func SendResponse(c *gin.Context, statusCode int, responseMessage string) {
	c.JSON(statusCode, gin.H{"messsage": responseMessage})
}
