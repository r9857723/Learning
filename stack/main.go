package main

import (
	"errors"
	"fmt"
)

/*
only + - * /
*/

type Stack struct {
	MaxTop int
	Top int
	array []int
}

func (s *Stack) Push(val int) (err error) {
	if s.MaxTop - 1 == s.Top {
		return errors.New("stack full")
	}
	s.Top++
	s.array = append(s.array, val)
	fmt.Println("Push success")
	return
}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		return 0, errors.New("stack empty")
	}
	val = s.array[s.Top]
	s.Top--
	return val , nil
}

func (s *Stack) ShowStack() {
	if s.Top == -1 {
		fmt.Println("Stack empty")
		return
	}
	//cTop = s.Top
	for i := s.Top; i >= 0 ; i-- {
		fmt.Printf("array[%d]=%d\n", i, s.array[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top: -1,
	}
	err := stack.Push(1)
	if err != nil {
		fmt.Println(err)
	}
	err = stack.Push(5)
	if err != nil {
		fmt.Println(err)
	}
	err = stack.Push(2)
	if err != nil {
		fmt.Println(err)
	}
	err = stack.Push(3)
	if err != nil {
		fmt.Println(err)
	}
	err = stack.Push(4)
	if err != nil {
		fmt.Println(err)
	}
	err = stack.Push(5)
	if err != nil {
		fmt.Println(err)
	}
	stack.ShowStack()

	val, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	val, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	stack.ShowStack()
}