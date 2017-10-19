package datatypes

import "fmt"

// Variable declared and default to 0
var z int

// Variable declared and assingn a default value
var y int = 1

// Multiple variable declaration

var (
	a = 1
	b = 2
	c = 3
)

func PrintVariables() {
	fmt.Println(a, b, c)
}
