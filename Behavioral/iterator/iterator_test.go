package iterator

import "fmt"

func ExampleIterator() {

	slice := []string{"a", "b", "c", "d"}
	container := &ConcreateContainer{slice}
	for i := container.GetIterator(); i.HasNext(); {
		v := i.Next()
		fmt.Println(v)
	}
	// Output:
	// a
	// b
	// c
	// d
}
