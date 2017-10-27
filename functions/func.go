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

	varArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1)

	closure()
}

func varArgs(a ...int) {
	for _, x := range a {
		fmt.Println("varArgs x:", x)
	}
}

func closure() {

	add := func(a, b int) (c int) {
		c = a + b
		return
	}
	fmt.Println(add(1, 2))
}
