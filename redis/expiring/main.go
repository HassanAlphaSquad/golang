package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
	defer client.Close()

	sessionKey := "session:username"

	// Check if the username is already in session (Redis)
	username, err := client.Get(ctx, sessionKey).Result()
	if err == redis.Nil {
		// If not found, ask for username
		fmt.Print("Enter your username: ")
		reader := bufio.NewReader(os.Stdin)
		username, _ = reader.ReadString('\n')
		username = username[:len(username)-1] // Remove newline

		// Store username in Redis for 20 seconds
		err := client.Set(ctx, sessionKey, username, 5*time.Second).Err()
		if err != nil {
			log.Fatal("Error storing session:", err)
		}
		fmt.Println("Username stored in cache for 3 seconds.")
	} else if err != nil {
		log.Fatal("Error retrieving session:", err)
	} else {
		fmt.Println("Retrieved username from cache:", username)
	}
}
