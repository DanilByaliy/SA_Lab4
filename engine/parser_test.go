package engine

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

func (s *MySuite) TestParse_1(c *C) {
	commandLine := "print hello"
	result := Parse(commandLine)
	expected1 := PrintCommand("hello")

	c.Assert(result, Equals, expected1)


	commandLine = "palindrom hello"
	result = Parse(commandLine)
	expected2 := PalindromCommand("hello")

	c.Assert(result, Equals, expected2)


	commandLine = "show hello"
	result = Parse(commandLine)
	expected1 = PrintCommand("SYNTAX ERROR: invalid command")

	c.Assert(result, Equals, expected1)


	commandLine = "     "
	result = Parse(commandLine)
	expected1 = PrintCommand("SYNTAX ERROR: no command")

	c.Assert(result, Equals, expected1)


	commandLine = "palindrom "
	result = Parse(commandLine)
	expected1 = PrintCommand("SYNTAX ERROR: no argument")

	c.Assert(result, Equals, expected1)
}