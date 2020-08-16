package main

type Strategy interface {
	Do()
}

type StrategyA struct {
	value float32
}

type StrategyB struct {
	value float64
}

func (strategy StrategyA) Do() {
	println(strategy.value)
}

func (strategy StrategyB) Do() {
	println(strategy.value)
}
