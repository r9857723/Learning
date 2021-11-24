package	main

import (
	"fmt"
)

type HeroNode struct {
	no int
	name string
	nickName string
	next *HeroNode
}

// InsertHeroNode insert heroNode to list
func (h *HeroNode) InsertHeroNode(newHeroNode *HeroNode) {
	tempNode := h
	for {
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next
	}
	tempNode.next = newHeroNode
}

// InsertHeroNodeAsc insert heroNode to list
func (h *HeroNode) InsertHeroNodeAsc(newHeroNode *HeroNode) {
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
		tempNode.next = newHeroNode
	}
}

func (h *HeroNode) DeleteHeroNode(id int) {
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

}

// ListHeroNode show all heroNode data
func (h *HeroNode) ListHeroNode() {
	tempNode := h
	if tempNode.next == nil {
		fmt.Println("HeroNode list is null")
		return
	}

	for {
		fmt.Printf("[%d, %s, %s] \n", tempNode.next.no, tempNode.next.name, tempNode.next.nickName)
		tempNode = tempNode.next
		if tempNode.next == nil {
			break
		}
	}
}

var heroHead = &HeroNode{}

func main() {
	hero3 := &HeroNode{
		no: 3,
		name: "Michael",
		nickName: "M",
	}
	heroHead.InsertHeroNodeAsc(hero3)

	hero1 := &HeroNode{
		no: 1,
		name: "Tom",
		nickName: "T",
	}
	heroHead.InsertHeroNodeAsc(hero1)

	hero2 := &HeroNode{
		no: 2,
		name: "Candy",
		nickName: "C",
	}
	heroHead.InsertHeroNodeAsc(hero2)
	heroHead.ListHeroNode()

	heroHead.DeleteHeroNode(2)
	heroHead.ListHeroNode()
}