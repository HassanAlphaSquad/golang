package main

import (
	"fmt"
	auth "golang/jwt"
)

func main() {
	// var input, upper_limit, start int
	// fmt.Print("Print a table of? ")
	// fmt.Scan(&input)
	// fmt.Print("Starting from? ")
	// fmt.Scan(&start)
	// fmt.Print("Upto? ")
	// fmt.Scan(&upper_limit)
	// for i := start; i <= upper_limit; i++ {
	// 	fmt.Printf("%d * %d = %d\n", input, i, i*input)
	// }
	jwtToken, err := auth.GenerateJWT("hassanisavailable@gmail.com", "123456789")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("JWT Token:", jwtToken)

	isValid := auth.JWTMiddlewareValidation(jwtToken)
	if isValid {
		fmt.Println("Token is valid")
	} else {
		fmt.Println("Token is invalid")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Lifetime")
	isAlive := auth.ValidateJWTToken(jwtToken)
	if isAlive {
		fmt.Println("Token is alive")
	} else {
		fmt.Println("Token is expired")
	}
}
