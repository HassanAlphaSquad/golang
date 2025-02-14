package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if username == "" {
		fmt.Println("Username cannot be empty!")
		return
	}

	keyPrefix := username + ":notes"
	client := redisClient()
	defer client.Close()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Note")
		fmt.Println("2. Retrieve Note")
		fmt.Println("3. Delete Note")
		fmt.Println("4. List Notes")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter note title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter note content: ")
			content, _ := reader.ReadString('\n')
			content = strings.TrimSpace(content)

			client.HSet(ctx, keyPrefix, title, content)
			fmt.Println("Note saved!")

		case "2":
			fmt.Print("Enter note title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			content, err := client.HGet(ctx, keyPrefix, title).Result()
			if err != nil {
				fmt.Println("Note not found!")
			} else {
				fmt.Println("Note Content:", content)
			}

		case "3":
			fmt.Print("Enter note title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			client.HDel(ctx, keyPrefix, title)
			fmt.Println("Note deleted!")

		case "4":
			notes, err := client.HKeys(ctx, keyPrefix).Result()
			if err != nil || len(notes) == 0 {
				fmt.Println("No notes found!")
			} else {
				fmt.Println("Your Notes:")
				for _, note := range notes {
					fmt.Println("-", note)
				}
			}

		case "5":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
