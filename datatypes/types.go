package datatypes

import "fmt"

var globalVariable string = "GlobalVariable"

func Print() {
	var a int
	var b int = 3

	fmt.Println("A ", a)
	fmt.Println("B ", b)

	fmt.Println("GlobalVariable", globalVariable)
}
