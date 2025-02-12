package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})
	healthCheck(client)

	// Set a key-value pair in Redis
	// setKey(client, "username", "hassanzahid")
	setKey(client, "country", "Pakistan")

	// Get the value from Redis
	// value := getKey(client, "username")
	country := getKey(client, "country")
	// fmt.Println("🔴 Value from Redis:", value)
	fmt.Println("🔴 Value from Redis: ", country)
}

// Function to check if Redis is running
func healthCheck(client *redis.Client) {
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}
	fmt.Println("✅ Redis Connected!", pong)
}

// Function to set a key in Redis
func setKey(client *redis.Client, key, value string) {
	err := client.Set(ctx, key, value, 1*time.Second).Err() // Expires in 10 minutes
	if err != nil {
		log.Fatal("Error setting key:", err)
	}
	fmt.Println("✅ Key set successfully!")
}

// Function to get a key from Redis
func getKey(client *redis.Client, key string) string {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("⚠️ Key does not exist")
		return ""
	} else if err != nil {
		log.Fatal("Error getting key:", err)
	}
	return val
}
