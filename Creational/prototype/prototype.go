package prototype

type Prototype interface {
	Clone()
}
type ConcretePrototype struct {
	name string
}

func (p *ConcretePrototype) Clone() *ConcretePrototype {
	return &ConcretePrototype{p.name}
}

func (p *ConcretePrototype) String() string {
	return "ConcretePrototype [name=" + p.name + "]"
}