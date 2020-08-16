package main

import (
	"fmt"
)

type Person struct {
	age  int8
	name string
}

type Clone interface {
	Clone() interface{}
}

func (person *Person) Clone() interface{} {
	return Person{person.age, person.name}
}

func main() {
	per1 := Person{1, "cxy"}
	per2 := per1.Clone()

	fmt.Println(per1)
	fmt.Println(per2)
	print(&per1.name)
	fmt.Println(&per1)
	fmt.Println(&per2)
	name := per2.(Person).name
	fmt.Println(&name)
}
