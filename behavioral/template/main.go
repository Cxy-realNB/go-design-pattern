package main

type Template interface {
	func1()
	func2()
}

func Do(template Template) {
	println(template)
	template.func1()
	template.func2()
}

type Actor struct {
}

func (Actor) func1() {}
func (Actor) func2() {}

func main() {
	actor := Actor{}
	println(actor)
	Do(actor)
}
