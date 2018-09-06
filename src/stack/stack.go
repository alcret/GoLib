package stack

import (
	"errors"
	"log"
	"testing"
)

type Stack struct {
	stack []int
}

func (s *Stack) CreateStack() {

}

func (s *Stack) PushStack(i int) {
	s.stack = append(s.stack, i)
}

func (s *Stack) PopStack() (int, error) {
	if len(s.stack) == 0 {
		return -1, errors.New("栈下溢出")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack) EmptyStack() bool {
	return len(s.stack) == 0
}

func TestStack(t *testing.T) {
	var stack Stack
	stack.PushStack(1)
	log.Println(stack.PopStack())
}
