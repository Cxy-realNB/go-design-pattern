package main

type Ints []int

func main() {
	ints := Ints{1, 2, 45, 85}
	//------------------------------------------
	Do(func(i int) {
		println(i)
	}, ints)
	//------------------------------------------
	it := ints.Iterate()
	for {
		val, ok := it()
		if !ok {
			break
		}
		println(val)
	}
}

// do implement
func Do(fn func(int), ints Ints) {
	for _, value := range ints {
		fn(value)
	}
}

// closure implement
func (ints Ints) Iterate() func() (int, bool) {
	index := 0
	return func() (i int, b bool) {
		if index >= len(ints) {
			return
		}
		i, ok := ints[index], true
		index++
		return i, ok
	}
}
