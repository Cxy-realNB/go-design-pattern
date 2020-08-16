package main

import "fmt"

type Fruit struct {
	color Color
}

type Apple struct {
	Fruit
	nativeTo string
}

type Grape struct {
	Fruit
	isCluster bool
}

type Introduce interface {
	Introduce()
}

func (fruit *Fruit) introduce() {
	fmt.Println(fruit)
}

func (apple *Apple) introduce() {
	fmt.Println(apple)
}

func (grape *Grape) introduce() {
	fmt.Println(grape)
}
