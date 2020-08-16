package main

type Visit interface {
	Visit(string)
}

type TestVisitor struct {
}

type ProdVisitor struct {
}

type Env struct {
	name  string
	visit Visit
}

func (env Env) Accept(visit Visit) {
	env.visit = visit
}

func (env Env) Visit() {
	env.visit.Visit(env.name)
}

func (TestVisitor) Visit(env string) {
	println("Test visitor visit " + env)
}

func (ProdVisitor) Visit(env string) {
	println("Prod visitor visit " + env)
}
