package main

func main() {
	redApple := &Apple{Fruit{Color{"red"}}, "NanTong"}
	redApple.introduce()
	greenGrape := &Grape{Fruit{Color{"green"}}, true}
	greenGrape.introduce()
}
