package observer

import "fmt"

//subscriber
type IObserver interface {
	Update(subj ISubject)
}

//publisher
type ISubject interface {
	AddObserver(o IObserver)
	RemoveObserver(o IObserver)
	NotifyObservers()
	GetInfo() string
}

//implements ISubject
type Event struct {
	eventInfo string
	observers []IObserver
}

func  NewEvent(info string) ISubject {
	ev := &Event{}
	ev.eventInfo = info
	return ev
}
func (ev *Event) GetInfo() string {
	return ev.eventInfo
}

func (ev *Event) AddObserver(o IObserver) {
	ev.observers = append(ev.observers, o)
}

func (ev *Event) RemoveObserver(o IObserver) {
	var indexToRemove int
	for i, observer := range ev.observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	ev.observers = append(ev.observers[:indexToRemove], ev.observers[indexToRemove+1:]...)
}

func (ev *Event) NotifyObservers() {
	for _, ob := range ev.observers {
		ob.Update(ev)
	}

}

//implements IObserver
type User struct{}

func (u *User) Update(subj ISubject) {
	info := subj.GetInfo()
	fmt.Println(info)
	/// do business logic with this info
}
