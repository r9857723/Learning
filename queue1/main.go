package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

// AddQueue 添加進隊列
func (q *Queue) AddQueue(val int) (err error) {
	// 判斷隊列是否已滿
	if q.rear == q.maxSize - 1 {
		return errors.New("Queue full")
	}
	q.rear++
	q.array[q.rear] = val
	return
}

// GetQueue 取出隊列中的數據
func (q *Queue) GetQueue() (val int, err error) {
	// 先判斷是否為空
	if q.front == q.rear {
		return -1, errors.New("Queue empty")
	}
	q.front++
	val = q.array[q.front]
	return val, err
}

// ShowQueue 顯示對列
func (q *Queue) ShowQueue() {
	fmt.Println("目前對列：")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, q.array[i])
	}
}

func main() {
	q := &Queue{
		maxSize: 5,
		rear: -1,
		front: -1,
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
			err := q.AddQueue(val)
			if err != nil {
				fmt.Println("加入隊列失敗")
			} else {
				fmt.Println("加入隊列成功")
			}
		case "get":
			fmt.Println("get")
			val, err := q.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("從隊列中取出了 %d \n", val)
			}
		case "show":
			fmt.Println("show")
			q.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}