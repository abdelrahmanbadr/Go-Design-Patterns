package adapter

func ExampleAdapter() {
	sparrow := Sparrow{}
	toyDuck := PlasticToyDuck{}
	// Wrap a bird in a birdAdapter so that it behaves like toy duck
	birdAdapter := BirdAdapter{&sparrow}
	sparrow.fly()
	sparrow.makeSound()
	toyDuck.squeak()
	birdAdapter.squeak()

	// Output:
	// 	Flying
	// Chirp Chirp
	// Squeak
	// Chirp Chirp

}
