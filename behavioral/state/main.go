package main

var state = true
var arr = [...]int{1, 2, 3, 4, 5}

func main() {
	Doing()
	SetState(true)
	println(state)
}

func SetState(flag bool) {
	state = flag
}

func Doing() {
	slice := arr[4:]
	println(slice[0])
}
