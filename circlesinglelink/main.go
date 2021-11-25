package main

import "fmt"

type HeroNode struct {
	no int
	name string
	next *HeroNode
}

// InsertHeroNode insert heroNode to link
func (h *HeroNode) InsertHeroNode(newHeroNode *HeroNode) {
	if h.next == nil {
		h.no = newHeroNode.no
		h.name = newHeroNode.name
		h.next = h
		fmt.Println("create first link")
		return
	}
	tempNode := h
	for {
		if tempNode.next == tempNode {
			break
		}
		tempNode = tempNode.next
	}
	tempNode.next= newHeroNode
	newHeroNode.next = h
}

// ListHeroNode show all heroNode data
func (h *HeroNode) ListHeroNode() {
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
	head.InsertHeroNode(hero1)
	hero2 := &HeroNode{
		no: 2,
		name: "Roman",
	}
	head.InsertHeroNode(hero2)
	head.ListHeroNode()
}