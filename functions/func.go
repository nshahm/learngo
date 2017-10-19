package functions

import "fmt"

func printNumber() {
	fmt.Println("Print number function invoked")
}

func functionWithSingleReturnValue() int {
	return 10
}

func functionWithMultipleReturnValue() (int, string) {
	return 10, "shahm"
}

func functionWithInputAndReturnValue(a int) int {
	return a
}

func PublicInvokefunction() {
	printNumber()
	a := functionWithSingleReturnValue()
	fmt.Println("A", a)

	id, name := functionWithMultipleReturnValue()
	fmt.Printf("id %v and name %v\n", id, name)
	// Function returns is added with blank identifier (_) and that is ignored
	_, name = functionWithMultipleReturnValue()
	fmt.Println("Name ", name)

	// Not necessary to declare and assigned value in two different step , better is :=
	var z int
	z = functionWithInputAndReturnValue(a)
	fmt.Println("Z ", z)
}
