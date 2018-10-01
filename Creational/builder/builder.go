package builder

import "fmt"

type Vehicle struct {
	Wheels    int
	Seats     int
	TopSpeed  int
	Structure string
}
type IVehilce interface {
	AcclerateSpeed(speed int) string
}

func (v Vehicle) AcclerateSpeed(speed int) {
	fmt.Println("speed of ", v.Structure, " is ", speed, "instead of ", v.TopSpeed)
}

type VehicleBuilder interface {
	SetWheels() VehicleBuilder
	SetSeats() VehicleBuilder
	SetStructure() VehicleBuilder
	TopSpeed() VehicleBuilder
	Build() Vehicle
}

//A Builder of type motorbike
type BikeBuilder struct {
	v Vehicle
}

func (b *BikeBuilder) SetWheels() VehicleBuilder {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() VehicleBuilder {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) TopSpeed() VehicleBuilder {
	b.v.TopSpeed = 180
	return b
}

func (b *BikeBuilder) SetStructure() VehicleBuilder {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) Build() Vehicle {
	return b.v
}

//A Builder of type Bus
type BusBuilder struct {
	v Vehicle
}

func (b *BusBuilder) SetWheels() VehicleBuilder {
	b.v.Wheels = 8
	return b
}
func (c *BusBuilder) SetSeats() VehicleBuilder {
	c.v.Seats = 52
	return c
}

func (b *BusBuilder) TopSpeed() VehicleBuilder {
	b.v.TopSpeed = 90
	return b
}

func (b *BusBuilder) SetStructure() VehicleBuilder {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) Build() Vehicle {
	return b.v
}

//A Builder of type Car
type CarBuilder struct {
	v Vehicle
}

func (c *CarBuilder) SetWheels() VehicleBuilder {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() VehicleBuilder {
	c.v.Seats = 5
	return c
}

func (b *CarBuilder) TopSpeed() VehicleBuilder {
	b.v.TopSpeed = 160
	return b
}

func (c *CarBuilder) SetStructure() VehicleBuilder {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() Vehicle {
	return c.v
}


type Director struct {
	vehicle VehicleBuilder
}
func (self *Director) Construct() {
	self.vehicle.SetSeats().SetStructure().SetWheels().TopSpeed()
}

func (self *Director) Build(b VehicleBuilder) Vehicle{
	self.vehicle = b
	self.Construct()
	return b.Build()
}