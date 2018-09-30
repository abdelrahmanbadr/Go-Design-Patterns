package strategy

type Strategy interface {
	Execute(int, int) int
}

type Add struct{}

func (self *Add) Execute(a, b int) int {
	return a + b
}

type Subtract struct{}

func (self *Subtract) Execute(a, b int) int {
	return a - b
}

type Multiply struct{}

func (self *Multiply) Execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (self *Context) SetStrategy(strategy Strategy) {
	self.strategy = strategy
}

func (self *Context) Calc(a, b int) int {
	return self.strategy.Execute(a, b)
}
