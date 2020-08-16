package main

type Proxy struct {
	actor Actor
}

type Actor struct {
	name string
}

type Action interface {
	Act()
}

func (proxy *Proxy) Act() {
	proxy.actor.Act()
}

func (proxy *Actor) Act() {
	print(proxy.name + " is doing")
}

func main() {
	proxy := &Proxy{Actor{"actor1"}}
	proxy.Act()
}
