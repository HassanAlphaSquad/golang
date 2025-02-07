package main

import "fmt"

func main() {
	fmt.Println("Struct in GO")
	type messageToSend struct {
		phoneNumber int
		address     string
	}
}
