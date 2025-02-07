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
	}

	fmt.Println("Original Array:", inputs)
	reversed_array := reverse(inputs)
	fmt.Print("Reversed Array: ")
	fmt.Print(reversed_array)
	fmt.Printf("\n")
}
