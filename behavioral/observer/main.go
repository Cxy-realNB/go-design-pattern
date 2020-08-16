package main

func main() {
	actor := &Actor{}
	observers := []Observer{{name: "observer1"}, {name: "observer2"}}
	actor.observers = observers
	actor.Act()
}
