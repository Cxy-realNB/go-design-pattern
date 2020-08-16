package main

import "strconv"

func main() {
	Do(1)
	Do(0)
	Do(-1)
}

func Do(i int) {
	switch {
	case i == 0:
		println("Value is 0")
	case i == -1:
		println("Value is -1")
	case i == 1:
		println("Value is 1")
	}
	println("print processed value is " + strconv.Itoa(i*i))
}
