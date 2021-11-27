package main

import (
	"fmt"
	"learning/expstack/stack"
	"strconv"
)

/* simple calculation */
func main() {
	numStack := &stack.Stack{
		MaxTop: 20,
		Top: -1	,
	}
	operStack := &stack.Stack{
		MaxTop: 20,
		Top: -1	,
	}

	exp := "3+2*6-2"
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	for {
		str := exp[index:index+1]
		strAscii := int([]byte(str)[0])
		fmt.Println(str, strAscii)
		if operStack.IsOperation(strAscii) {
			// operaStack is empty
			if operStack.Top == -1 {
				operStack.Push(strAscii)
			} else {
				resPrev, _ := operStack.Priority(operStack.Array[operStack.Top])
				resNow, _ := operStack.Priority(strAscii)
				if resPrev >= resNow {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()

					result, _ = operStack.Cal(num1, num2, oper)

					numStack.Push(result)
				}
				operStack.Push(strAscii)
			}
		} else {
			val, _ := strconv.ParseInt(str, 10, 64)
			numStack.Push(int(val))
		}
		index ++
		if index > len(exp) - 1 {
			break
		}
	}

	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()

		result, _ = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}

	number, _ := numStack.Pop()
	fmt.Println(number)

}

