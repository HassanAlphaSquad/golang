package main

import "fmt"

func main() {
	car := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	colours := map[string]int{"Red": 1, "Blue": 2, "Black": 3}
	subjects := map[int]string{1: "Mathematics", 2: "Biology", 3: "Chemistry", 4: "Physics"}
	fmt.Printf("%v", car)
	fmt.Printf("\n")
	fmt.Printf("%v", colours)
	fmt.Printf("\n")
	// for key, value := range colours {
	// 	fmt.Println(key, "=>", value)
	// }
	// fmt.Print(colours)
	for key, value := range subjects {
		fmt.Println(key, "=>", value)
	}
	fmt.Printf("\n")

}
