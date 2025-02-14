package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Args:")

	if len(os.Args) < 4 {
		fmt.Println("Put 3 arguments after your go run command:")
		fmt.Println("1st: Name, 2nd: Age, 3rd: Profession")
		fmt.Println("More args expected")
		return
	}
	// os.Args[0] is for program's name
	name := os.Args[1]
	age := os.Args[2]
	profession := os.Args[3]

	fmt.Printf("Your name is: %v -- Your age is: %v years -- Your profession is: %v\n", name, age, profession)
}
