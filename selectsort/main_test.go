package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_selectSort1(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	Sorting(arr)

	t.Log("Success")
}

func Test_selectSort2(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	Sorting(arr)

	t.Log("Success")
}

func Test_selectSort3(t *testing.T) {
	var arr []int
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 200000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	Sorting(arr)

	t.Log("Success")
}