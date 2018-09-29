package command

func ExampleCommand() {
	light := &Light{}
	lightOff := NewLightOffCommand(light)
	lightOn := NewLightOnCommand(light)

	control := new(RemoteControl)
	control.setCommand(lightOn)
	control.pressButton()

	control.setCommand(lightOff)
	control.pressButton()

	// Output:
	// light is on
	// light is off

}
