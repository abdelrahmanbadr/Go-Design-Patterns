package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
type ConcreateContainer struct {
	Items []string
}

func (c *ConcreateContainer) GetIterator() Iterator {
	return &ConcreateIterator{Container: c}
}
func (c *ConcreateContainer) AddItem(s string) {
	c.Items = append(c.Items, s)
}

type ConcreateIterator struct {
	Container    *ConcreateContainer
	currentIndex int
}

func (c *ConcreateIterator) HasNext() bool {
	if c.currentIndex < len(c.Container.Items) {
		return true
	}
	return false
}

func (c *ConcreateIterator) Next() interface{} {
	if c.HasNext() != true {
		return nil
	}
	defer func() {
		c.currentIndex++
	}()
	return c.Container.Items[c.currentIndex]
}

// or Aggregate
type Container interface {
	GetIterator() Iterator
}
