package main

type Handle interface {
	Handle()
	Next() *Handler
}

type Handler struct {
	next *Handler
	name string
}

func (handler *Handler) Handle() {
	println(handler.name)
}

func (handler *Handler) Next() *Handler {
	return handler.next
}
