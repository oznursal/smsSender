package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"smsSender/model"
	"time"
)

var ctx = context.Background()

func connectRedis() (*redis.Client, error) {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password for local development
		DB:       0,                // Default DB
	})

	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {

		}
	}(rdb)

	// Ping the Redis server to check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
		log.Fatal("Error connecting to Redis:", err)
	}

	fmt.Println("Connected to Redis:", pong)
	return rdb, err
}

func SaveSms(response *model.SmsResponse) error {
	key := response.MessageId
	value := time.Now().Format("2006-01-02 15:04:05")
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {

		}
	}(rdb)

	rdb.Set(ctx, key, value, 0)

	return nil
}
