package example1

import (
	"fmt"
)

// Component
type Beverage interface {
	Description() string
}

// ConcreteComponent
type BasicBeverage struct{}

func (c *BasicBeverage) Description() string {
	return ""
}

// Concrete Decorator
type Coca struct {
	beverage Beverage
}

func (c *Coca) Description() string {
	if c.beverage == nil {
		return "Coco"
	} else {
		return c.beverage.Description() + " Coco"
	}
}

// Concrete Decorator
type Coffe struct {
	beverage Beverage
}

func (c *Coffe) Description() string {
	if c.beverage == nil {
		return "Coffe"
	} else {
		return c.beverage.Description() + " Coffe"
	}
}

func main() {
	component := &Coffe{&Coca{new(BasicBeverage)}}
	fmt.Println(component.Description())
}
