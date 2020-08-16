package main

type Command interface {
	Execute()
}

type CommandA struct {
}

type CommandB struct {
}

func (command CommandA) Execute() {
	println("commandA executing")
}

func (command CommandB) Execute() {
	println("commandB executing")
}
