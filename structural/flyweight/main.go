package main

type SharingObject struct{}
type Component struct{ SharingObject }

func main() {
	so := SharingObject{}
	c1 := &Component{so}
	c2 := &Component{so}
	println(c1)
	println(c2)
}
