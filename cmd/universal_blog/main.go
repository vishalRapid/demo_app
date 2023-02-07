package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/vishalrana9915/demo_app/pkg/algolia"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/ratelimiter"
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

	var mongoURI string = os.Getenv("MONGO_URI")
	var mongoDB string = os.Getenv("MONGO_DB")

	redisConnector.ConnectToRedis(redis_url, pass)

	databaseConnector.Adapter.Connect(mongoURI, mongoDB)

	// setup search
	algolia.Adapter.SetupAlgolia()

	router := gin.Default()

	rateLimit := ratelimiter.SetupLimiter()

	// Create a new middleware with the limiter instance.
	middleware := mgin.NewMiddleware(rateLimit)
	// setting up the limiter
	router.Use(middleware)

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
