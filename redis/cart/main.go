package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// userID = "hassan"
)

func main() {
	defer client.Close()
	var userID string
	fmt.Print("Enter your username 👤: ")
	fmt.Scanf("%s", &userID)
	fmt.Print("\n🛒 Welcome to the E-commerce Cart System!\n\n")
	for {
		fmt.Print("🎲 Choices:\n\n")
		fmt.Println("[1] Add Item ➕")
		fmt.Println("[2] Remove Item 🗑️")
		fmt.Println("[3] Update Quantity 🔄")
		fmt.Println("[4] View Cart 🛒")
		fmt.Println("[5] Clear Cart 🧹")
		fmt.Println("[6] Exit 🚪")

		switch input("\nEnter choice: ") {
		case "1":
			add_item(userID)
		case "2":
			remove_item(userID)
		case "3":
			update_quantity(userID)
		case "4":
			view_cart(userID)
		case "5":
			clear_cart(userID)
		case "6":
			fmt.Println("\n👋 Exiting. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("\n⛔ Invalid choice. Try again.")
		}
	}
}

func add_item(userID string) {
	item := input("\nEnter item name: ")
	qty_str := input("Enter quantity: ")
	quantity, err := strconv.Atoi(qty_str)
	if err != nil || quantity <= 0 {
		fmt.Println("\n⛔ Invalid quantity.")
		return
	}

	cart_key := "cart:" + userID                                             //  cart:userID(hassan)
	if err := client.HSet(ctx, cart_key, item, quantity).Err(); err != nil { // apple 10 => HSET cart:hassan apple 10 (redis-cli)
		log.Fatalf("\n⛔ Error adding item: %v", err)
	}
	fmt.Println("\n✅ Item added to cart!")
	seperator(15)
}

func remove_item(userID string) {
	item := input("\nEnter item name to remove: ")
	cart_key := "cart:" + userID

	if _, err := client.HDel(ctx, cart_key, item).Result(); err != nil {
		fmt.Println("\n⛔ Error removing item:", err)
	} else {
		fmt.Println("\n✅ Item removed from cart!")
		seperator(15)
	}
}

func update_quantity(userID string) {
	item := input("\nEnter item name: ")
	qty_str := input("Enter new quantity: ")
	quantity, err := strconv.Atoi(qty_str)
	if err != nil || quantity <= 0 {
		fmt.Println("\n⛔ Invalid quantity.")
		return
	}

	cart_key := "cart:" + userID
	if err := client.HSet(ctx, cart_key, item, quantity).Err(); err != nil {
		log.Fatalf("\n⛔ Error updating quantity: %v", err)
	}
	fmt.Println("\n✅ Quantity updated!")
	seperator(15)
}

func view_cart(userID string) {
	cart_key := "cart:" + userID
	items, err := client.HGetAll(ctx, cart_key).Result()
	if err != nil {
		fmt.Println("\n⛔ Error fetching cart:", err)
		return
	}

	if len(items) == 0 {
		fmt.Print("\n🛒 Your cart is empty.\n\n")
		return
	}
	fmt.Println("\n👤 User ID: ", userID)
	fmt.Print("\n🛍️  Your Cart:\n\n")
	for item, qty := range items {
		fmt.Printf("  ⮕  %-8s: %-8s\n", item, qty)
	}
	fmt.Printf("\n📋 Total items in cart: %d\n\n", len(items))
	seperator(25)
}

func clear_cart(userID string) {
	cart_key := "cart:" + userID
	if _, err := client.Del(ctx, cart_key).Result(); err != nil {
		fmt.Println("\n⛔ Error clearing cart:", err)
	} else {
		fmt.Println("\n✅ Cart cleared!")
		seperator(10)
	}
}

func input(prompt string) string {
	fmt.Print(prompt)
	var text string
	fmt.Scanln(&text)
	return text
}

func seperator(size int) {
	fmt.Println("\n" + strings.Repeat("*", size) + "\n")
}
