package main

import "fmt"

func reverse(array [5]int) [5]int {
	var reversed [5]int
	n := len(array)

	for i := 0; i < n; i++ {
		reversed[i] = array[n-1-i]
	}

	return reversed
}
func print_array(array [5]int) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d->", array[i])
	}
	print("\n")
}

func main() {
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)
	var inputs [5]int
	for i := 0; i < len(inputs); i++ {
		fmt.Printf("Enter value%d: ", i+1)
		fmt.Scan(&inputs[i])
		// fmt.Print("\n")
	}

	// fmt.Println("User input:", inputs)
	print_array(inputs)
	reversed := reverse(inputs)
	fmt.Print(reversed)
}
