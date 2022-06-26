package main

import (
	"bufio"
	engine "lab4/engine"
	"os"
)

func main() {
	inputFile := "commands.txt"
	eventLoop := new(engine.Loop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
   		defer input.Close()
   		scanner := bufio.NewScanner(input)
   		for scanner.Scan() {
     			commandLine := scanner.Text()
        	cmd := engine.Parse(commandLine) // parse the line to get a Command 
        	eventLoop.Post(cmd)
   		}
	}
	eventLoop.AwaitFinish()
}