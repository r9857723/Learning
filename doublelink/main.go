package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickName string
	prev     *HeroNode // previous node
	next     *HeroNode // next node
}

// InsertHeroNode insert heroNode to list
func InsertHeroNode(h *HeroNode, newHeroNode *HeroNode) {
	tempNode := h
	for {
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next
	}
	tempNode.next = newHeroNode
	newHeroNode.prev = tempNode
}

// InsertHeroNodeAsc insert heroNode to list
func InsertHeroNodeAsc(h *HeroNode, newHeroNode *HeroNode) {
	tempNode := h
	flag := true
	for {
		if tempNode.next == nil {
			break
		} else if tempNode.next.no > newHeroNode.no {
			break
		} else if tempNode.next.no == newHeroNode.no {
			flag = false
			break
		}
		tempNode = tempNode.next
	}
	if !flag {
		fmt.Println("編號重複 no =", newHeroNode.no)
		return
	} else {
		newHeroNode.next = tempNode.next
		newHeroNode.prev = tempNode
		if tempNode.next != nil {
			tempNode.next.prev = newHeroNode
		}
		tempNode.next = newHeroNode
	}
}

func DeleteHeroNode(h *HeroNode, id int) {
	tempNode := h

	flag := false
	for {
		if tempNode.next == nil {
			break
		} else if tempNode.next.no == id {
			flag = true
			break
		}
		tempNode = tempNode.next
	}
	if !flag {
		fmt.Println("id not exist")
		return
	}
	tempNode.next = tempNode.next.next
	if tempNode.next != nil {
		tempNode.next.prev = tempNode
	}
}

// ListHeroNode show all heroNode data
func ListHeroNode(h *HeroNode) {
	tempNode := h
	if tempNode.next == nil {
		fmt.Println("HeroNode list is null")
		return
	}

	for {
		fmt.Printf("[%d, %s, %s] \n", tempNode.next.no, tempNode.next.name, tempNode.next.nickName)
		//fmt.Println(tempNode.next)
		tempNode = tempNode.next
		if tempNode.next == nil {
			break
		}
	}
}

// ListHeroNodeDesc show all heroNode data
func ListHeroNodeDesc(h *HeroNode) {
	tempNode := h

	if tempNode.next == nil {
		fmt.Println("HeroNode list is null")
		return
	}

	//get last
	for {
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next
	}

	for {
		fmt.Printf("[%d, %s, %s] \n", tempNode.no, tempNode.name, tempNode.nickName)
		//fmt.Println(tempNode)
		tempNode = tempNode.prev
		if tempNode.prev == nil {
			break
		}
	}
}

var heroHead = &HeroNode{}

func main() {

	hero1 := &HeroNode{
		no:       1,
		name:     "Tom",
		nickName: "T",
	}
	InsertHeroNode(heroHead, hero1)

	hero2 := &HeroNode{
		no:       2,
		name:     "Candy",
		nickName: "C",
	}
	InsertHeroNode(heroHead, hero2)

	hero3 := &HeroNode{
		no:       3,
		name:     "Michael",
		nickName: "M",
	}
	InsertHeroNode(heroHead, hero3)
	ListHeroNode(heroHead)

	fmt.Println()
	ListHeroNodeDesc(heroHead)
}
