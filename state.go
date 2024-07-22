package goldmarkdocx

import "github.com/yuin/goldmark/ast"

type state struct {
	containerType ast.NodeKind

	textStyle      Style
	firstParagraph bool
	listType       listType
	itemNumber     int
	isHeading      bool
	isFooter       bool
	isHeader       bool
}

type states struct {
	stack []*state
}

func newStates() *states {
	return &states{
		stack: make([]*state, 0),
	}
}

func (s *states) push(c *state) {
	s.stack = append(s.stack, c)
}

func (s *states) pop() *state {
	var x *state
	x, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return x
}

func (s *states) peek() *state {
	return s.stack[len(s.stack)-1]
}

func (s *states) parent() *state {
	return s.stack[len(s.stack)-2]
}
