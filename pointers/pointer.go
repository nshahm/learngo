package pointers

import (
	"fmt"
)

type Vehicle struct {
	Name string
	VIN  int
}

func NewVehicle(name string, vin int) *Vehicle {
	return &Vehicle{Name: name, VIN: vin}
}

func PrintVehicle(v *Vehicle) {
	fmt.Printf("Vehile is %v with VIN %v\n", v.Name, v.VIN)
}
