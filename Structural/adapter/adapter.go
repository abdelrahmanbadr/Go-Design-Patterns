package adapter

import "fmt"

// Adaptee defines the old interface which needs to be adapted.
type Bird interface {
	fly()
	makeSound()
}
type Sparrow struct{}

func (this *Sparrow) fly() {
	fmt.Println("Flying")
}
func (this *Sparrow) makeSound() {
	fmt.Println("Chirp Chirp")
}

// Target is the target interface to adapt to.
type ToyDuck interface {
	squeak()
}
type PlasticToyDuck struct{}

func (this *PlasticToyDuck) squeak() {
	fmt.Println("Squeak")
}

// Adapter is the adaptor to the new interface Target.
type BirdAdapter struct {
	bird Bird
}

func (this *BirdAdapter) squeak() {
	this.bird.makeSound();
}
