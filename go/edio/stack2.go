package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Stack is empty")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	stk := *stack
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*stack = stk[:len(stk)-1]
	return top, nil
}

var st1 Stack

func main() {
	st1.Push("Brown")
	st1.Push(3.14)
	st1.Push(100)
	st1.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
