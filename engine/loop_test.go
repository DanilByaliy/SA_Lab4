package engine

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestLoop(c *C) {
	command := PrintCommand("first print")
	command2 := PalindromCommand("palind")
	loop := new(Loop)
	loop.Start()

	result1 := loop.stop
	result2 := len(loop.q.a)
	c.Assert(result1, Equals, false)
	c.Assert(result2, Equals, 0)


	loop.Post(command)
	loop.Post(command2)

	result1 = loop.stop
	result2 = len(loop.q.a)
	c.Assert(result1, Equals, false)
	c.Assert(result2, Equals, 2)


	loop.AwaitFinish()

	result1 = loop.stop
	result2 = len(loop.q.a)
	c.Assert(result1, Equals, true)
	c.Assert(result2, Equals, 0)


	loop.Post(command)

	result1 = loop.stop
	result2 = len(loop.q.a)
	c.Assert(result1, Equals, true)
	c.Assert(result2, Equals, 0) // The command is not 
	//added to the event loop queue because the user 
	// stopped receiving it with the "AwaitFinish()" command

}
