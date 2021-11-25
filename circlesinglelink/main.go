package main

import "fmt"

type HeroNode struct {
	no int
	name string
	next *HeroNode
}

// InsertHeroNode insert heroNode to link
func InsertHeroNode(h *HeroNode, newHeroNode *HeroNode) {
	if h.next == nil {
		h.no = newHeroNode.no
		h.name = newHeroNode.name
		h.next = h
		fmt.Println("create first link")
		return
	}
	tempNode := h
	for {
		if tempNode.next == h {
			break
		}
		tempNode = tempNode.next
	}
	tempNode.next= newHeroNode
	newHeroNode.next = h
}

func DeleteHeroNode(h *HeroNode, no int) *HeroNode {

	tempNode := h
	helperNode := h
	//tempNode := h

	// 空鏈表
	if tempNode.next == nil {
		fmt.Println("link is null")
		return h
	}
	// 只有頭節點
	if tempNode.next == h {
		if tempNode.no == no {
			tempNode.next = nil
		}
		return h
	}

	for {
		if helperNode.next == h {
			break
		}
		helperNode = helperNode.next
	}

	flag := true
	for {
		if tempNode.next == h {
			break
		}
		if tempNode.no == no {
			if tempNode == h {
				h = h.next
			}
			helperNode.next = tempNode.next
			flag = false
			break
		}
		tempNode = tempNode.next
		helperNode = helperNode.next
	}
	if flag {
		if tempNode.no == no {
			helperNode.next = tempNode.next
		} else {
			fmt.Println("no not exist")
		}
	}
	return h
}

// ListHeroNode show all heroNode data
func ListHeroNode(h *HeroNode) {

	tempNode := h
	if tempNode.next == nil {
		fmt.Println("link is null")
		return
	}
	for {
		fmt.Println(tempNode)
		if tempNode.next == h {
			break
		}
		tempNode = tempNode.next
	}
}


func main() {
	// init head
	head := &HeroNode{}
	hero1 := &HeroNode{
		no: 1,
		name: "Tom",
		next: nil,
	}
	InsertHeroNode(head, hero1)
	hero2 := &HeroNode{
		no: 2,
		name: "Roman",
	}
	InsertHeroNode(head, hero2)
	hero3 := &HeroNode{
		no: 3,
		name: "Yuki",
	}
	InsertHeroNode(head, hero3)
	ListHeroNode(head)

	head = DeleteHeroNode(head, 1)
	//head.DeleteHeroNode(2)
	//head.DeleteHeroNode(3)
	ListHeroNode(head)
}