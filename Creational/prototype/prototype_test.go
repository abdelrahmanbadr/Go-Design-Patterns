package prototype

import "fmt"

func ExamplePrototype(){
	prototype := &ConcretePrototype{"prototype1"}
	fmt.Println(prototype.name)
	fmt.Println(prototype.Clone().name)

	//Output
	//prototype1
	//prototype1
}
