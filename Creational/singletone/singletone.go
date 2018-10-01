package singletone

type singletone struct {
	counter int
}
type SingletoneContract interface {
	IncrementCounter() int
}

var object *singletone

func GetSingletoneObject() *singletone {
	if object == nil {
		object = new(singletone)
	}
	return object
}
func (s *singletone) IncrementCounter() int {
	s.counter++
	return s.counter
}