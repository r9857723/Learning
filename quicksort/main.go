package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//array := []int{9,-3,7,99,5,4,-99,2,-1}
	////array := []int {99, -1, -7}
	//quickSort(0, len(array) - 1, array)
	//fmt.Println("main", array)

	var arr []int
	start := time.Now().Unix()
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10000000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	end := time.Now().Unix()
	fmt.Printf("create arr %d s \n", end-start)

	start = time.Now().Unix()
	quickSort(0, len(arr) -1, arr)
	end = time.Now().Unix()
	fmt.Printf("sorting time %d s \n", end-start)
}

func quickSort(left, right int, array []int) {
	//fmt.Println("input = ", array)
	l := left
	r := right
	pivot := array[(l + r) / 2]
	//fmt.Printf("pvoit = %d \n", pivot)
	for ; l < r; {
		for ; array[l] < pivot ; {
			l++
		}
		for ; array[r] > pivot ; {
			r--
		}
		if l >= r{
			//fmt.Println("B")
			break
		}
		//fmt.Printf("%d \t與 %d\t交換 \t= ", array[l], array[r])
		array[r], array[l] = array[l], array[r]
		//fmt.Println("~=", array)
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l ++
		r --
	}
	//fmt.Println("output = ", array)

	if left < r {
		quickSort(left, r, array)
	}

	if right > l {
		quickSort(l, right, array)
	}
}