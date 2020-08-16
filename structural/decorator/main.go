package main

import "fmt"

type Decoratee struct {
}

type Decorator struct {
	decoratee *Decoratee
}

type Do interface {
	Do()
}

func (Decoratee) Do() {
	fmt.Println("I am good!")
}

func (decorator Decorator) Do() {
	fmt.Println("adornment ...")
	decorator.decoratee.Do()
	fmt.Println("adornment ...")
}

func main() {
	decorator := &Decorator{&Decoratee{}}
	decorator.Do()
}
