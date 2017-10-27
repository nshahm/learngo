package builtintypes

import "fmt"

func DeclareArray() {
	var arr [2]int

	//arr = make([2]int)
	arr[0] = 1
	arr[1] = 2

	fmt.Println("Declared array and assigned manually", arr)
	fmt.Println("Length of the declared array", len(arr))
}

func InitializeArray() {

	var arr1 = [3]int{1, 2, 3}
	fmt.Println("Array declard and initialized", arr1)
	initializeUsingMake()
}

func printArray(arr []int) {

	for _, i := range arr {
		fmt.Println("Array print using for loop arr x", i)
	}
}

func initializeUsingMake() {

	arr2 := make([]int, 10)

	var i int
	for i = 1; i < 9; i++ {
		arr2[i-1] = int(i)
	}

	printArray(arr2)
}
