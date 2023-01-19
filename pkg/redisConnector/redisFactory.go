package redisConnector

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/matryer/try"
)

var RedisClient *redis.Client

/*
*
We need this function to make connection to redis, this function
will take two params
One for redis url to connect to
*/
func ConnectToRedis(url string, redisPass string) {
	// Wait for Redis to start up
	err := try.Do(func(attempt int) (bool, error) {
		var err error
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     url,
			Password: redisPass,
			DB:       0,
		})
		_, err = RedisClient.Ping().Result()
		if err != nil {
			log.Println("Error connecting to Redis, retrying...")
			time.Sleep(10 * time.Second)
			return attempt < 5, err
		}
		return true, nil
	})

	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
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

// set value for a hash in redis
func HSetValue(hash string, key string, value string) {
	err := RedisClient.HSet(hash, key, value).Err()
	if err != nil {
		panic(err)
	}
}

// fetch key from hget
func HGetValue(hash string, key string) string {
	fmt.Println(hash, key, RedisClient)
	result, err := RedisClient.HGet(hash, key).Result()

	if err != nil {
		panic(err)
	}

	return result
}

// Checking if user exist in Caching or not
func CheckHExist(hash string) bool {
	// Check if a hash exists
	exists, err := RedisClient.Exists(hash).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)

	}

	// returns boolean, 1 means user exist and 0 means it doesn't
	if exists > 0 {
		return true
	} else {
		return false
	}
}
