package error

import "fmt"

func InvokePanic() {

	defer func() {
		err := recover()
		fmt.Println("Error recovered ", err)
	}()
	x, y := 1, 0
	a := x / y
	fmt.Println(a)
	panic("Panic Invoked")
}
