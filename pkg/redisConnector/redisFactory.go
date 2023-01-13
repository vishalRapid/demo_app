package redisConnector

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

/*
*
We need this function to make connection to redis, this function
will take two params
One for redis url to connect to
*/
func ConnectToRedis(url string, redisPass string) {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: redisPass,
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Error making connection to redis", err)
	}

	fmt.Println("Cache connected..")
}

// setting value to the redis store
func GetValue(key string) (string, error) {
	return RedisClient.Get(key).Result()
}

// fetching value from the store
func SetValue(key string, value string) error {
	return RedisClient.Set(key, value, 0).Err()
}
