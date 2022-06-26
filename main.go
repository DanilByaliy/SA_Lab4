package main

import (
	engine "lab4/engine"
)

func main() {
	loop := new(engine.Loop)

	loop.Start()

	//
	loop.Post(engine.PrintCommand("hello"))
	//
	loop.Post(engine.PalindromCommand("hello"))
	//
	loop.Post(engine.PrintCommand("hello2"))

	loop.AwaitFinish()

	loop.Post(engine.PrintCommand("hello2"))
}
