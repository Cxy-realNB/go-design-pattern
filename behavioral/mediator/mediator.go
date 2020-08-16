package main

type Register interface {
	Register()
}

type Sender struct {
	message string
}

type Receiver struct {
	message string
}

type Mediator struct {
	sender   *Sender
	receiver *Receiver
}

func (mediator Mediator) Send() {
	mediator.receiver.message = mediator.sender.message
	println(mediator.receiver.message)
}

func (mediator Mediator) Register(sender *Sender, receiver *Receiver) {
	mediator.sender = sender
	mediator.receiver = receiver
}
