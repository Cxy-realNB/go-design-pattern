package main

func main() {
	handler1 := Handler{name: "handle1"}
	handler2 := Handler{name: "handle2"}
	handler1.next = &handler2
	handler1.Handle()
	handler1.next.Handle()
}
