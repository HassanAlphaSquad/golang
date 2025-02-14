package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Redis client setup
func redis_client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   15,
	})
}

// Subscribe to chat messages
func subscribe(channel, username string, quit chan bool) {
	client := redis_client()
	defer client.Close()

	pub_sub := client.Subscribe(ctx, channel)
	defer pub_sub.Close()

	// Notify others about the new user
	joining_msg := fmt.Sprintf("📢 %s has joined the chat!", username)
	client.Publish(ctx, channel, joining_msg)

	for {
		select {
		case <-quit:
			return // Stop listening when user quits
		case msg := <-pub_sub.Channel():
			if msg == nil || msg.Payload == "" {
				continue // Ignore empty messages
			}

			sender, message := parse_msg(msg.Payload)

			// Ignore self-join notification
			if msg.Payload == joining_msg {
				continue
			}

			// Ignore self messages
			if sender == username {
				continue
			}

			// Display messages
			if sender == "SYSTEM" {
				fmt.Printf("\n%s\n", message)
			} else {
				fmt.Printf("\n%s: %s\n", sender, message)
			}

			// Print user prompt again
			fmt.Print(username + "# ")
		}
	}
}

// Publish messages to chat
func publish(channel, username string, quit chan bool) {
	client := redis_client()
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(username + "# ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		// Ignore empty messages
		if message == "" {
			continue
		}

		// Handle exit command
		if message == "/exit" {
			client.Publish(ctx, channel, fmt.Sprintf("SYSTEM: 🚪 %s has left the chat.", username))
			fmt.Println("\nExiting chat...")
			quit <- true
			return
		}

		// Publish message in "username: message" format
		client.Publish(ctx, channel, fmt.Sprintf("%s: %s", username, message))
	}
}

// Detects abrupt disconnection
func handle_disconnect(username, channel string) {
	client := redis_client()
	defer client.Close()

	for {
		time.Sleep(5 * time.Second) // Check every 5 seconds

		// Try to PING Redis, if it fails, user is disconnected
		_, err := client.Ping(ctx).Result()
		if err != nil {
			client.Publish(ctx, channel, fmt.Sprintf("SYSTEM: ❌ %s has unexpectedly disconnected!", username))
			fmt.Println("\n❌ Connection lost. Exiting chat...")
			os.Exit(1)
		}
	}
}

// Parse message format ("username: message")
func parse_msg(msg string) (string, string) {
	parts := strings.SplitN(msg, ": ", 2)
	if len(parts) < 2 {
		return "SYSTEM", msg // Handle system messages
	}
	return parts[0], parts[1]
}

func main() {
	// Ask for username
	fmt.Print("Enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if username == "" {
		fmt.Println("Username cannot be empty!")
		return
	}

	// Default chat channel
	channel := "global_chat"

	// Quit channel to stop goroutines
	quit := make(chan bool)

	// Handle graceful exit (CTRL+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\n🚪 Closing chat...")
		quit <- true
		os.Exit(0)
	}()

	// Start listening for messages
	go subscribe(channel, username, quit)

	// Detect abrupt disconnection
	go handle_disconnect(username, channel)

	// Delay before publishing to allow subscription
	time.Sleep(1 * time.Second)

	// Start sending messages
	publish(channel, username, quit)
}
