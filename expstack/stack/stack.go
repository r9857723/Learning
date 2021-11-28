package stack


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
	Array [5]int
}

func (s *Stack) Push(val int) (err error) {
	if s.MaxTop - 1 == s.Top {
		return errors.New("stack full")
	}
	s.Top++
	s.Array[s.Top] = val
	return
}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		return 0, errors.New("stack empty")
	}
	val = s.Array[s.Top]
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
		fmt.Printf("array[%d]=%d\n", i, s.Array[i])
	}
}

func (s *Stack) IsOperation(val int) bool {
	switch val {
		case 43, 45, 42, 47:
			return true
	}
	return false
}

func (s *Stack) Cal(n1, n2 int, operation int) (result int, err error) {
	switch operation {
		case 43:
			result = n2 + n1
		case 45:
			result = n2 - n1
		case 42:
			result = n2 * n1
		case 47:
			result = n2 / n1
		default:
			err = errors.New("operation err")
	}
	return
}

func (s *Stack) Priority(operation int) (result int, err error) {
	switch operation {
		case 42, 47:
			result = 1
		case 43, 45:
			result = 0
		default:
			err = errors.New("operation err")
	}
	return
}
