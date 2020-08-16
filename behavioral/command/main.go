package main

type Receiver struct {
	command Command
}

func main() {
	receiver1 := &Receiver{command: CommandA{}}
	receiver1.command.Execute()

	receiver2 := &Receiver{command: CommandB{}}
	receiver2.command.Execute()
}
