package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 34, 19, 100, 80}
	var arr []int
	start := time.Now().Unix()
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 200000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	end := time.Now().Unix()
	fmt.Printf("create arr %d s \n", end-start)

	start = time.Now().Unix()
	Sorting(arr)
	end = time.Now().Unix()
	fmt.Printf("sorting time %d s \n", end-start)
	//fmt.Println(arr)
	//sorting2(arr)
	//fmt.Println(arr)
}

// Sorting 由大到小
func Sorting(arr []int) {
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
}

// Sorting2 由小到大
func Sorting2(arr []int) {
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
}