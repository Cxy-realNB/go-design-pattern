package main

type Component struct {
	composite1 Composite1
	composite2 Composite2
}

type Composite1 struct {
}

type Composite2 struct {
}

func (composite1 *Composite1) Composite1Do() {

}

func (composite2 *Composite2) Composite2Do() {

}

func main() {
	component := &Component{composite1: Composite1{}, composite2: Composite2{}}
	component.composite1.Composite1Do()
	component.composite2.Composite2Do()
}
