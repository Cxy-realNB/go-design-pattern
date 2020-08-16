package main

import "strconv"

var previous int
var now int

func main() {
	Do(1)
	Do(2)
}

func Do(i int) {
	previous = now
	now = i
	println("previous:" + strconv.Itoa(previous) + " now:" + strconv.Itoa(now))
}
