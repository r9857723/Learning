package main

import "fmt"

type node struct {
	no int
	next *node
}



func CreateNodes(num int) *node {
	firstNode := &node{}
	tempNode := &node{}
	if num < 1 {
		return firstNode
	}

	for i := 1; i <= num; i++ {
		boy := &node{
			no: i,
		}
		if i == 1 {
			firstNode = boy
		} else {
			tempNode.next = boy
		}
		tempNode = boy
		tempNode.next = firstNode
	}
	return firstNode
}

func ShowNode(first *node) {
	if first.next == nil {
		fmt.Println("first is null")
		return
	}
	tempNode := first
	for {
		fmt.Println(tempNode)
		if tempNode.next == first {
			break
		}
		tempNode = tempNode.next
	}
}

func Count(first *node) int {
	if first.next == nil {
		return 0
	}
	tempNode := first
	var cnt int
	for {
		cnt++
		if tempNode.next == first {
			break
		}
		tempNode = tempNode.next
	}
	return cnt
}

func DoJosephus(h *node, startNum int, countNum int) {
	if h.next == nil {
		fmt.Println("link is not exist")
		return
	}

	// check limit
	if Count(h) < startNum {
		fmt.Println("out of index")
		return
	}

	//get last
	tailNode := h
	for {
		if tailNode.next == h {
			break 
		}
		tailNode = tailNode.next
	}

	// move to startNum	
	for i := 0; i < startNum -1; i++ {
		h = h.next
		tailNode = tailNode.next
	}
	
	for {
		for i := 1; i <= countNum -1; i++ {
			h = h.next
			tailNode = tailNode.next
		}


		fmt.Print(h)
		fmt.Println("Out")
		h = h.next
		tailNode.next = h

		if h == tailNode {
			break
		}
	}
	fmt.Print(h)
	fmt.Println("Out")

}

func main() {
	nodes := CreateNodes(5)
	ShowNode(nodes)
	fmt.Println()
	DoJosephus(nodes, 2, 3)
}
