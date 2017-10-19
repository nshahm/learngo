package arithmetic

import "fmt"

// Arithmetic type
type Arithmetic struct{}

/*
Add function
*/
func (ar *Arithmetic) Add(a, b int) int {
	c := a + b
	fmt.Println("The value of the C is ", c)
	return c
}
