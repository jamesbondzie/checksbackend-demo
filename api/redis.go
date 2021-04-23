package api

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)


var client *redis.Client





//RedisServer function to:
//1.load env variables
//2.connect to redis
//3.connect/initialize postgres DB
func RedisServer() {
	
	var err error

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("$REDIS_URL must be set")
	}

	client  =  redis.NewClient(&redis.Options{
		Addr : redisURL, //redis port
	})

	_, err = client.Ping().Result()
	if err != nil {
	   panic(err)
	}

	fmt.Println("REDIS_DNS connected")


}