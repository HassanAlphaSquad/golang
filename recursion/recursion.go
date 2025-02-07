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
	fac20 := Factorial(20)
	fmt.Print(fac20)
	fmt.Print("\n")
}
