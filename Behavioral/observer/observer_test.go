package observer

func ExampleObserver() {
	user := &User{}
	ev := NewEvent("Event Info")
	ev.AddObserver(user)
	ev.NotifyObservers()

	// Output:
	// Event Info

}
