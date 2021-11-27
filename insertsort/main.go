package main

import "fmt"

func main() {
	arr := []int{23, 0, 12, 56, 34}
	insertSort(arr)
}

// 由大到小
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertValue := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] < insertValue {
			arr[insertIndex + 1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertValue
		}
	}

	fmt.Println(arr)
}