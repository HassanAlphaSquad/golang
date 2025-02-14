package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func redis_client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func shorten_URL(url string, expiry int) {
	client := redis_client()
	defer client.Close()
	// code := fmt.Sprintf("url:%d", time.Now().Unix()) // if multiple users request for url-shortening with at same time, it will generate same keys
	code := fmt.Sprintf("%d", time.Now().UnixNano()) // generates code which represents number of nano-seconds elapsed from epoch time (using nano keeping in view the case when multiple users are trying to shorten url)
	client.Set(ctx, code, url, time.Duration(expiry)*time.Second)
	fmt.Println("Shortened URL:", code)
}

func get_URL(code string) {
	client := redis_client()
	defer client.Close()
	url, err := client.Get(ctx, code).Result()
	if err != nil {
		fmt.Println("URL not found or expired.")
		return
	}
	fmt.Println("Original URL:", url)
}

func ttl(code string) {
	client := redis_client()
	defer client.Close()

	duration, err := client.TTL(ctx, code).Result()
	if err != nil {
		fmt.Println("Error fetching TTL:", err)
		return
	}

	if duration == -2 {
		fmt.Println("Error: Shortened URL not found or expired.")
		return
	}

	fmt.Printf("Time-to-live (TTL) for %s: %v seconds\n", code, int(duration.Seconds()))
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  shorten <url> [expiry in seconds (OPTIONAL, DEFAULT: 1 hour)]")
		fmt.Println("  get <short-code>")
		fmt.Println("  ttl <short-code>")
		return
	}

	// os.Args[0] -> main.go (program's name)
	// os.Args[1] -> shorten / get / ttl
	// os.Args[2] -> url / short-code
	// os.Args[3] -> expiry time (seconds)

	switch os.Args[1] {
	case "shorten":
		expiry := 3600
		if len(os.Args) > 3 {
			_, err := fmt.Sscanf(os.Args[3], "%d", &expiry) // taking Args[3] from user and storing in expiry in decimal form
			if err != nil {
				fmt.Println("Invalid expiry value. Using default 3600 seconds.")
			}
		}
		shorten_URL(os.Args[2], expiry)
	case "get":
		get_URL(os.Args[2])
	case "ttl":
		ttl(os.Args[2])
	default:
		fmt.Println("Unknown command.")
	}
}
