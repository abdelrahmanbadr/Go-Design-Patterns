package facade

func ExampleFacade() {
	computer := NewComputer()
	computer.Start()

	// Output:
	// 	CPU frozen
	// HardDrive reads
	// Memory loads
	// CPU jumps to 0
	// CPU executes

}
