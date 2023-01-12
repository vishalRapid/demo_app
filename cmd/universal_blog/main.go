package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishalrana9915/demo_app/pkg/redisConnector"
	"github.com/vishalrana9915/demo_app/pkg/responseHandler"
	"github.com/vishalrana9915/demo_app/pkg/routes"
)

func ReturnUsers(c *gin.Context) {

	responseHandler.SendResponse(c, 200, "success")
}

func HandlerError(c *gin.Context) {
	// c.JSON(400, c.H{"message": userDummyData})
	responseHandler.SendResponse(c, 404, "Work under progress")
}

func main() {

	godotenv.Load()

	var redis_url string = os.Getenv("REDIS_URL")
	var pass string = os.Getenv("REDIS_PASS")

	fmt.Println("Running in debugging mode ==>", gin.IsDebugging())
	redisConnector.ConnectToRedis(redis_url, pass)

	router := gin.Default()

	port := os.Getenv("SERVER_PORT")

	// setting up routes
	routes.SetupRouter(router)

	// checking if port exist in env or not
	if port == "" {
		port = "8080"
		fmt.Print("No port found in .env")
	}

	router.Run(":" + port)

}
