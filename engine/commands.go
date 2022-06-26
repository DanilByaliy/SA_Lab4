package engine

import (
	"fmt"
	"strings"
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

func (pa PalindromCommand) Execute(h Handler) {
	 s1 := strings.Split(string(pa), "")
  	 s2 := make([]string, 0)
  	 for i:=len(s1)-1; i>=0; i-- {
    	 	s2 = append(s2, s1[i])
  	 }
  	 ap := strings.Join(s2, "")
  	 ap = string(pa) + ap

  	 h.Post(PrintCommand(ap))
}