package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
	sessionDuration = 20 * time.Second // Session expires in 20 seconds
)

func main() {
	defer client.Close()

	fmt.Println("🚀 Welcome to Console-Based Session Management")

	sessionID, err := client.Get(ctx, "session_id").Result()
	if err == redis.Nil {
		// No session found, ask for username
		username := getUsername()
		sessionID = uuid.NewString()

		// Store session in Redis
		err := client.Set(ctx, "session_id", sessionID, sessionDuration).Err()
		if err != nil {
			log.Fatal("Error storing session ID:", err)
		}

		err = client.Set(ctx, "session:"+sessionID, username, sessionDuration).Err()
		if err != nil {
			log.Fatal("Error storing username:", err)
		}

		fmt.Println("✅ Session created successfully!")
	} else {
		fmt.Println("✅ Existing session found. Retrieving user info...")
		username, err := client.Get(ctx, "session:"+sessionID).Result()
		if err == redis.Nil {
			fmt.Println("❌ Session expired. Please restart.")
			return
		} else if err != nil {
			log.Fatal("Error retrieving session:", err)
		}
		fmt.Printf("🔑 Welcome back, %s!\n", username)
	}

	// Show options
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Show session info")
		fmt.Println("2. Logout")
		fmt.Println("3. Exit")

		choice := getInput("Enter choice: ")

		switch choice {
		case "1":
			showSessionInfo()
		case "2":
			logout()
			return
		case "3":
			fmt.Println("👋 Exiting. Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}

// Function to get username input
func getUsername() string {
	username := getInput("Enter your username: ")
	return username
}

// Function to retrieve and show session info
func showSessionInfo() {
	sessionID, _ := client.Get(ctx, "session_id").Result()
	username, err := client.Get(ctx, "session:"+sessionID).Result()
	if err == redis.Nil {
		fmt.Println("❌ No active session found.")
	} else if err != nil {
		fmt.Println("❌ Error retrieving session:", err)
	} else {
		fmt.Printf("🔑 Logged in as: %s (Session ID: %s)\n", username, sessionID)
	}
}

// Function to handle logout
func logout() {
	sessionID, _ := client.Get(ctx, "session_id").Result()

	// Delete session from Redis
	client.Del(ctx, "session_id")
	client.Del(ctx, "session:"+sessionID)

	fmt.Println("✅ You have been logged out.")
	main()
}

// Function to get input from the user
func getInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1] // Remove newline character
}
