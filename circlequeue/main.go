package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array [5]int
	head int
	tail int
}

func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		return errors.New("Queue full")
	}
	cq.array[cq.tail] = val
	cq.tail = (cq.tail + 1) % cq.maxSize
	return
}

func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		return 0, errors.New("Queue empty")
	}
	val = cq.array[cq.head]
	cq.head= (cq.head + 1) % cq.maxSize
	return
}

func (cq *CircleQueue) ListQueue() {
	size := cq.Size()
	if size == 0 {
		fmt.Println("Queue is empty")
		return
	}

	tempHead := cq.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d \t", tempHead, cq.array[tempHead])
		tempHead = (tempHead + 1) % cq.maxSize
	}
	fmt.Println()
}

func (cq *CircleQueue) IsFull() bool {
	return (cq.tail + 1) % cq.maxSize == cq.head
}

func (cq *CircleQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

func (cq *CircleQueue) Size() int {
	return (cq.tail + cq.maxSize - cq.head) % cq.maxSize
}

func main() {

	q := &CircleQueue{
		maxSize: 5,
		tail: 0,
		head: 0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 輸入add 表示添加到對列")
		fmt.Println("2. 輸入get 表示從對列獲取數據")
		fmt.Println("3. 輸入show 表示顯示隊列")
		fmt.Println("4. 輸入exit 離開")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Printf("請輸入需添加的數字")
			fmt.Scanln(&val)
			err := q.Push(val)
			if err != nil {
				fmt.Println("加入隊列失敗")
			} else {
				fmt.Println("加入隊列成功")
			}
		case "get":
			val, err := q.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("從隊列中取出了 %d \n", val)
			}
		case "show":
			q.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}