package main

import "fmt"

func main() {
	var input, upper_limit, start int
	fmt.Print("Print a table of? ")
	fmt.Scan(&input)
	fmt.Print("Starting from? ")
	fmt.Scan(&start)
	fmt.Print("Upto? ")
	fmt.Scan(&upper_limit)
	for i := start; i <= upper_limit; i++ {
		fmt.Printf("%d * %d = %d\n", input, i, i*input)
	}
}
