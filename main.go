package main

import (
	"bufio"
	"fmt"
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
    			fmt.Println(commandLine)
        		// cmd := parse(commandLine) // parse the line to get a Command 
        		// eventLoop.Post(cmd)
   		}
	}
	eventLoop.AwaitFinish()
}