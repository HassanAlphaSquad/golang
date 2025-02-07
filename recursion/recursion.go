package main

import "fmt"

func Factorial(x float64) (y float64) {
	if x > 0 {
		y = x * Factorial(x-1)
	} else {
		y = 1
	}
	return y
}

func main() {
	fmt.Print("Factorial:")
	fmt.Print(Factorial(10))
	fmt.Print("\n")
}
