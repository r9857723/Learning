package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 34, 19, 100, 80}

	rtn := sorting(a)
	fmt.Println(rtn)
	rtn2 := sorting2(a)
	fmt.Println(rtn2)
}


func sorting(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
		}
	}
	return arr
}


func sorting2(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
	return arr
}