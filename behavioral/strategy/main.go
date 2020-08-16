package main

type Actor struct {
	strategy Strategy
}

func main() {
	actor := new(Actor)
	actor.strategy = &StrategyA{1.0}
	actor.strategy.Do()
}
