package main

type Actor struct {
	observers []Observer
}

type Observer struct {
	name string
}

func (actor Actor) Act() {
	println("I am working!")
	for _, value := range actor.observers {
		value.Observe()
	}
}

func (observer Observer) Observe() {
	println("I got you! I am " + observer.name)
}
