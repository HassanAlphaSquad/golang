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
	ctx              = context.Background()
	client           = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	session_duration = 20 * time.Second
)

func main() {
	defer client.Close()
	fmt.Print("\n🚀 Welcome to Console-Based Session Management\n")

	sessionID, err := client.Get(ctx, "session_id").Result()
	if err == redis.Nil {
		new_session()
	} else {
		resume_session(sessionID)
	}
	menu()
}

func new_session() {
	username := input("\nEnter your username: ")
	sessionID := uuid.NewString()
	if err := client.Set(ctx, "session_id", sessionID, session_duration).Err(); err != nil {
		log.Fatal("Error storing session ID:", err)
	}
	if err := client.Set(ctx, "session:"+sessionID, username, session_duration).Err(); err != nil {
		log.Fatal("Error storing username:", err)
	}
	fmt.Println("\n✅ Session created successfully!")
}

func resume_session(sessionID string) {
	username, err := client.Get(ctx, "session:"+sessionID).Result()
	if err == redis.Nil {
		fmt.Println("\n❌ Session expired. Please restart.")
		os.Exit(0)
	} else if err != nil {
		log.Fatal("\nError retrieving session:", err)
	}
	fmt.Printf("\n🔑 Welcome back, %s!\n", username)
}

func menu() {
	for {
		fmt.Println("\n1. Show session info\n2. Logout\n3. Exit")
		switch input("\nEnter choice: ") {
		case "1":
			session_info()
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

func session_info() {
	sessionID, err := client.Get(ctx, "session_id").Result()
	if err == redis.Nil {
		fmt.Println("❌ No active session found.")
		return
	} else if err != nil {
		fmt.Println("❌ Error retrieving session ID:", err)
		return
	}

	username, err := client.Get(ctx, "session:"+sessionID).Result()
	if err == redis.Nil {
		fmt.Println("❌ Session expired or invalid.")
	} else if err != nil {
		fmt.Println("❌ Error retrieving session:", err)
	} else {
		fmt.Printf("\n🔑 Logged in as: %s (Session ID: %s)\n", username, sessionID)
	}
}

func logout() {
	sessionID, _ := client.Get(ctx, "session_id").Result()
	client.Del(ctx, "session_id", "session:"+sessionID)
	fmt.Println("\n✅ You have been logged out.")

	for {
		fmt.Println("\n1. Start program again? ")
		fmt.Println("2. Exit program?")
		choice := input("\nEnter choice: ")

		switch choice {
		case "1":
			fmt.Print("\n==============================================\n")
			main()
		case "2":
			fmt.Println("\n👋 Exiting. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}

func input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}
