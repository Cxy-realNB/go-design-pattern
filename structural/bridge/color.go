package main

type Color struct {
	name string
}

type Red struct {
	Color
}

type Green struct {
	Color
}

type GetName interface {
	GetName()
}

func (red Red) GetName() string {
	print("red")
	return "red"
}

func (green Green) GetName() string {
	print("green")
	return "green"
}
