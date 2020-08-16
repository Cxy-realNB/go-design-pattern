package main

func main() {
	env := Env{name: "Cob", visit: TestVisitor{}}
	env.Visit()
}
