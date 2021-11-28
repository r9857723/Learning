package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//arr := []int{23, 0, 12, 56, 34}
	//insertSort(arr)


	var arr []int
	start := time.Now().Unix()
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 200000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	end := time.Now().Unix()
	fmt.Printf("create arr %d s \n", end-start)

	start = time.Now().Unix()
	InsertSort(arr)
	end = time.Now().Unix()
	fmt.Printf("sorting time %d s \n", end-start)
}

// 由大到小
func InsertSort(arr []int) {
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

	//fmt.Println(arr)
}