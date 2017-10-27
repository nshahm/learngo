package builtintypes

import "fmt"

func InitSlice() {
	var num = make([]int, 5, 5)

	// Slice is used to mention the length and the capacity not like a fixed length (array)

	//num = make([]int, 2, 5)
	fmt.Printf("Slice length = %d , capacity : %d, slice is %v ", len(num), cap(num), num)
	//num = [1 2]

	// Use slice methods

	num = append(num, 0, 1, 2, 3, 4, 5)
	fmt.Println("Slice 0: ", num)

	var slice1 = make([]int, 2, 4)
	slice1 = append(slice1, 6, 7, 8, 9)
	fmt.Println("Slice 1: ", slice1)

	var num2 = copy(num, slice1)
	fmt.Println("Copy slices: ", num2)

	fmt.Println("After copy slice :", num)
}
