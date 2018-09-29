package main

//Command
type Command interface {
	Execute()
}

//Concrete Command
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}
func (lightOn *LightOnCommand) Execute() {
	lightOn.light.switchOn()
}

//Concrete Command
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}
func (lightOfF *LightOffCommand) Execute() {
	lightOfF.light.switchOff()
}

// Reciever
type Light struct {
	lightOn bool
}

func (light *Light) switchOn() {
	light.lightOn = true
}
func (light *Light) switchOff() {
	light.lightOn = false
}

//Invoker
type RemoteControl struct {
	command Command
}

func (this *RemoteControl) setCommand(command Command) {
	this.command = command
}
func (this *RemoteControl) pressButton() {
	this.command.Execute()
}

//Client
func main() {
	control := new(RemoteControl)
	light := new(Light)
	lightsOn := NewLightOnCommand(light)
	lightsOff := NewLightOffCommand(light)

	//switch on
	control.setCommand(lightsOn)
	control.pressButton()

	//switch off
	control.setCommand(lightsOff)
	control.pressButton()

}
