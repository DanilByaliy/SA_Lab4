package engine

import (
	"fmt"
	"sync"
)

type commandsQueue struct {
	mu 	 sync.Mutex
	a 	 []Command
	wait bool

	notEmpty chan struct{}
}

func (cq *commandsQueue) push(c Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.a = append(cq.a, c)
	if cq.wait {
		cq.notEmpty <-struct{}{}
	}
}

func (cq *commandsQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.a) == 0 {
		cq.wait = true
		cq.mu.Unlock()
		<- cq.notEmpty
		cq.mu.Lock()
	} 

	res := cq.a[0]
	cq.a[0] = nil
	cq.a = cq.a[1:]
	return res
}

func (cq *commandsQueue) empty() bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.a) == 0
}

type Loop struct {
	q *commandsQueue

	stop 			 bool
	stopSignal chan struct{}
}

func (l *Loop) Start() {
	l.q = &commandsQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *Loop) Post(cmd Command) {
	// TODO: що робити, коли команда додана після 
	//       того як запит на зупинку було зроблено
	if l.stop == true {
		fmt.Println(cmd)
	}
	l.q.push(cmd)
}

type CommandFunc func(h Handler)

func (f CommandFunc) Execute(h Handler) {
	f(h)
}

func (l *Loop) AwaitFinish() {
	l.Post(CommandFunc(func(h Handler) {
		l.stop = true
	}))
	<- l.stopSignal
}
