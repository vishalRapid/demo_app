package redisConnector

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

/*
*
We need this function to make connection to redis, this function
will take two params
One for redis url to connect to
*/
func ConnectToRedis(url string, redisPass string) {

	client = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: redisPass,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Error making connection to redis", err)
	}

	fmt.Println("Database connected..")
}

// setting value to the redis store
func GetValue(key string) (string, error) {
	return client.Get(key).Result()
}

// fetching value from the store
func SetValue(key string, value string) error {
	return client.Set(key, value, 0).Err()
}
