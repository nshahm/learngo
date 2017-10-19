package main

import (
	"fmt"
	"time"

	"github.com/nshahm/learngo/arithmetic"
	"github.com/nshahm/learngo/datatypes"
	"github.com/nshahm/learngo/functions"
	"github.com/nshahm/learngo/pointers"
)

func main() {

	fmt.Println("hello world")
	fmt.Println("current time", time.Now())
	operation := &arithmetic.Arithmetic{}
	c := operation.Add(1, 2)
	fmt.Println("Added value", c)

	// datatypes
	datatypes.Print()
	datatypes.PrintVariables()
	// Pointers
	vehicle := pointers.NewVehicle("Honda civic", 1234567890)
	pointers.PrintVehicle(vehicle)

	// Creating instance using new
	vehicle1 := new(pointers.Vehicle)
	vehicle1.Name = "Toyota sienna"
	vehicle1.VIN = 9876543210

	pointers.PrintVehicle(vehicle1)

	// Functions
	functions.PublicInvokefunction()

}
