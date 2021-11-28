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
	operaStack := &stack.Stack{
		MaxTop: 20,
		Top: -1	,
	}

	exp := "30+30*6-4-6"
	index := 0
	num1 := 0
	num2 := 0
	opera := 0
	result := 0
	keepNum := ""
	for {
		str := exp[index:index+1]
		strAscii := int([]byte(str)[0])
		fmt.Println(str, strAscii)
		if operaStack.IsOperation(strAscii) {
			// operaStack is empty
			if operaStack.Top == -1 {
				operaStack.Push(strAscii)
			} else {
				resPrev, _ := operaStack.Priority(operaStack.Array[operaStack.Top])
				resNow, _ := operaStack.Priority(strAscii)
				if resPrev >= resNow {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					opera, _ = operaStack.Pop()

					result, _ = operaStack.Cal(num1, num2, opera)

					numStack.Push(result)
				}
				operaStack.Push(strAscii)
			}
		} else {

			keepNum += str
			if index == len(exp) - 1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operaStack.IsOperation(int([]byte(exp[index+1:index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
			//val, _ := strconv.ParseInt(str, 10, 64)
			//numStack.Push(int(val))
		}
		index ++
		if index > len(exp) - 1 {
			break
		}
	}

	fmt.Println(numStack)
	fmt.Println(operaStack)
	for {
		if operaStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		opera, _ = operaStack.Pop()

		result, _ = operaStack.Cal(num1, num2, opera)
		numStack.Push(result)
	}

	number, _ := numStack.Pop()
	fmt.Println(number)

}

