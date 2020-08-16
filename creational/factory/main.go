package main

type Factory struct {
}

func (factory Factory) ProductJuice() string {
	return "juice"
}

func (factory Factory) ProductMilk() string {
	return "milk"
}

type Product interface {
	ProductMilk() string
	ProductJuice() string
}

var fac Factory = Factory{}

func main() {
	milk := fac.ProductMilk()
	juice := fac.ProductJuice()
	print(milk)
	print(juice)
}
