package singletone

import "fmt"

func ExampleObserver() {
	object1 := GetSingletoneObject()
	fmt.Println(object1.IncrementCounter())
	object2 := GetSingletoneObject()
	fmt.Println(object2.IncrementCounter())
	object3 := GetSingletoneObject()
	fmt.Println(object3.IncrementCounter())


	// Output:
	// 1
	// 2
	// 3

}
