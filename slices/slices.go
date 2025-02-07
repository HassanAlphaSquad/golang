package main

import "fmt"

func main() {
	slice := []int{32, 36, 56, 88, 99, 65, 45}
	fmt.Println("Original Slice:")
	fmt.Println("Last value of slice ->", slice[len(slice)-1])
	fmt.Println("Length of slice ->", len(slice))
	fmt.Println("Total capacity of slice ->", cap(slice))
	fmt.Println("Slice =>", slice)
	slice = append(slice, 77)
	fmt.Printf("Updated Slice => %d, Updated length => %d, Updated capacity => %d, Updated last index => %d\n\n", slice, len(slice), cap(slice), slice[len(slice)-1])

	dup_slice := make([]int, len(slice))
	copy(dup_slice, slice)
	fmt.Println("Duplicate Slice:")
	fmt.Println("Last value of duplicate slice ->", dup_slice[len(dup_slice)-1])
	fmt.Println("Length of duplicateslice ->", len(dup_slice))
	fmt.Println("Total capacity of duplicate slice ->", cap(dup_slice))
	fmt.Println("Duplicate Slice =>", dup_slice)
	dup_slice = append(dup_slice, 77)
	fmt.Printf("Updated duplicate slice => %d, Updated duplicate slice length => %d, Updated duplicate slice capacity => %d, Updated duplicate slice last index => %d\n\n", dup_slice, len(dup_slice), cap(dup_slice), dup_slice[len(dup_slice)-1])

	working_days := make([]string, 5)
	working_days[0] = "Monday"
	working_days[1] = "Tuesday"
	working_days[2] = "Wednesday"
	working_days[3] = "Thursday"
	working_days[4] = "Friday"
	fmt.Println("Working days =>", working_days)
}
