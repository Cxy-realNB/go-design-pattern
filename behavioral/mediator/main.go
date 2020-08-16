package main

func main() {
	mediator := &Mediator{}
	sender := &Sender{message: "ok"}
	receiver := &Receiver{}
	mediator.Register(sender, receiver)
	mediator.Send()
}
