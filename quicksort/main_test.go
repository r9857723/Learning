package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_quickSort(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	quickSort(0, len(arr) -1, arr)

	t.Log("Success")
}

func Test_quickSort2(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	quickSort(0, len(arr) -1, arr)

	t.Log("Success")
}

func Test_quickSort3(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	quickSort(0, len(arr) -1, arr)

	t.Log("Success")
}

func Test_quickSort4(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10000000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	quickSort(0, len(arr) -1, arr)

	t.Log("Success")
}

func Test_quickSort5(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 100000000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	quickSort(0, len(arr) -1, arr)

	t.Log("Success")
}