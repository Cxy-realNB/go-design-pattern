package main

import "fmt"

type Person struct {
	age  int8
	name string
}

type Builder struct {
	age  int8
	name string
}

func (builder Builder) Build(age int8, name string) Person {
	return Person{age, name}
}

func main() {
	builder := new(Builder)
	person := builder.Build(1, "CXY")
	fmt.Println(person)
}
