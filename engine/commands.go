package engine

import (
	"fmt"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type PrintCommand string

func (pc PrintCommand) Execute(h Handler) {
	 fmt.Println(string(pc))
}

type PalindromCommand string

func (mc PalindromCommand) Execute(h Handler) {
	ap := ""
	// TODO: implement the function
	h.Post(PrintCommand(ap))
}