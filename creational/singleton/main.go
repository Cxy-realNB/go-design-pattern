package main

import "sync"

type Singleton struct{}

type AddressPrinter interface {
	PrintAddress()
}

var singleton *Singleton
var once sync.Once

func (singleton *Singleton) PrintAddress() {
	print(singleton)
}

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()
	instance1.PrintAddress()
	instance2.PrintAddress()
}
