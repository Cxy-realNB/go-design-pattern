package main

type Target interface {
	Request()
}

type Adapter struct {
	Adaptor
}

type Adaptor struct {
}

func (adaptor *Adaptor) request() {
	print("doing")
}

func main() {
	adapter := &Adapter{Adaptor{}}
	adapter.request()
}
