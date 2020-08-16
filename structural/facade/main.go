package main

type Facade1 struct{}
type Facade2 struct{}
type Facade3 struct{}

func (Facade1) Do() {

}

func (Facade2) Do() {

}

func (Facade3) Do() {

}

func main() {
	facade1 := &Facade1{}
	facade1.Do()
	facade2 := &Facade2{}
	facade2.Do()
	facade3 := &Facade3{}
	facade3.Do()
}
