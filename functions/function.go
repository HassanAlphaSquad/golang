package main

import "fmt"

func table(number int, start int, end int) (status bool) {
	fmt.Print("\n")

	defer fmt.Println("Table function called!")
	if start > end {
		fmt.Printf("Invalid range!\n")
		status = false
		return
	}

	for i := start; i <= end; i++ {
		mul := start * i
		fmt.Printf("%d * %d = %d\n", number, i, mul)
	}
	fmt.Print("\n")
	status = true
	return
}

func divide(dividend int64, divisor int64) (quotient int64, reminder int64) {
	if divisor == 0 {
		fmt.Println("Dividing by 0 attempted")
		return 0, 0
	}
	quotient = dividend / divisor
	reminder = dividend % divisor
	return quotient, reminder
}

func add(nums ...int) {
	fmt.Println("Number(s) => ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Their sum => ", total)
}

func main() {
	table(2, 1, 10)
	table(10, 10, 1)
	q, r := divide(10, 20)
	fmt.Printf("Q=>%d, R=>%d\n\n", q, r)
	add(10, 10, 10, 10, 10)
	add(10, 20)
	add(1, 1)
}
