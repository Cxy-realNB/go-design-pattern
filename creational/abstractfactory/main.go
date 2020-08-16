package main

type Factory interface {
	ProdMilk() string
	ProdJuice() string
}

type Factory1 struct{}
type Factory2 struct{}

func (fac Factory1) ProdMilk() string {
	return "milk1"
}

func (fac Factory2) ProdMilk() string {
	return "milk2"
}

func (fac Factory1) ProdJuice() string {
	return "juice1"
}

func (fac Factory2) ProdJuice() string {
	return "juice2"
}

func main() {
	fac1 := Factory1{}
	fac2 := Factory2{}
	fac1.ProdJuice()
	fac2.ProdMilk()
}
