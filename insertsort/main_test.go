package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_insertSort1(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	InsertSort(arr)

	t.Log("Success")
}

func Test_insertSort2(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	InsertSort(arr)

	t.Log("Success")
}

func Test_insertSort3(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 200000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	InsertSort(arr)

	t.Log("Success")
}