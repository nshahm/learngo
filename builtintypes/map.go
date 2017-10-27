package builtintypes

import "fmt"

func Maps() {
	//var x map[int]string

	x := make(map[int]string)
	x[1] = "Shahm"
	x[2] = "Ghazni"
	x[3] = "x"

	fmt.Println("Map x : ", x)

	// Delete a value in map using key
	delete(x, 3)
	fmt.Println("After deleted 3 : ", x)

	// Search for non-exist key and return empty result

	fmt.Println("Non exist key x[4]", x[4])

	if name, ok := x[4]; ok {
		fmt.Println(name, ok)
	}

	if name, ok := x[2]; ok {
		fmt.Println(name, ok)
	}
}
