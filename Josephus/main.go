package main

import "fmt"

type Boy struct {
	no int
	next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 {
			first = boy
		} else {
			curBoy.next = boy
		}
		curBoy = boy
		curBoy.next = first
	}
	return first
}

func ShowBoy(first *Boy) {
	if first.next == nil {
		fmt.Println("first is null")
		return
	}
	tempBoy := first
	for {
		fmt.Println(tempBoy)
		if tempBoy.next == first {
			break
		}
		tempBoy = tempBoy.next
	}
}

func main() {
	link := AddBoy(5)

	ShowBoy(link)
}
