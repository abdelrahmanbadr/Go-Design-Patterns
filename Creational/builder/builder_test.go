package builder

import "fmt"

func ExampleBuilder() {
	director := Director{}
	vehicle := director.Build(&BikeBuilder{})
	fmt.Println(vehicle.TopSpeed)
	fmt.Println(vehicle.Structure)
	fmt.Println(vehicle.Seats)
	fmt.Println(vehicle.Wheels)

	// Output:
	// 180
	// Motorbike
	// 2
	// 2

}
