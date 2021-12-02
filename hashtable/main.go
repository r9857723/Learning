package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id int
	Name string
	Next *Emp
}

func (e *Emp) Show() {
	fmt.Printf("id=%d, name=%s \n", e.Id, e.Name)
}

type Emplink struct {
	Head *Emp
}

func (e *Emplink) Insert(emp *Emp) {
	temp := e.Head
	var pre *Emp = nil
	// 空鏈表
	if temp == nil {
		e.Head = emp
		return
	}
	for {
		if temp != nil {
			if temp.Id > emp.Id {
				break
			}
			pre = temp
			temp = temp.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = temp
}

func (e *Emplink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("link %d is empty \n", no)
		return
	}
	temp := e.Head
	for {
		fmt.Printf(" link %d  id=%d, name=%s \n", no, temp.Id, temp.Name)
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	fmt.Println()
}

func (e *Emplink) FindById(id int) *Emp {
	temp := e.Head
	for {
		if temp != nil && temp.Id == id {
			return temp
		} else if temp == nil {
			break
		}
		temp = temp.Next
	}
	return nil
}

type HashTable struct {
	Links [7]Emplink
}

func (h *HashTable) Insert(emp *Emp) {
	linkNo := h.HashFun(emp.Id)
	h.Links[linkNo].Insert(emp)
}

func (h *HashTable) HashFun(id int) int {
	return id % 7
}

func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.Links); i++ {
		h.Links[i].ShowLink(i)
	}
}

func (h *HashTable) FindById(id int) *Emp {
	linkNo := h.HashFun(id)
	return h.Links[linkNo].FindById(id)
}

func main () {
	var hashTable HashTable
	key := ""
	id := 0
	name := ""
	for {
		fmt.Println("input enter 1")
		fmt.Println("show enter 2")
		fmt.Println("find enter 3")
		fmt.Println("exit enter 4")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("id=")
			fmt.Scanln(&id)
			fmt.Println("name=")
			fmt.Scanln(&name)
			emp := &Emp{
				Id: id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "2":
			fmt.Println(2)
			hashTable.ShowAll()
		case "3":
			fmt.Println(3)
			fmt.Println("id=")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id = %d is not exist \n", id)
			} else {
				emp.Show()
			}

		case "4":
			os.Exit(0)
		default:
			fmt.Println()
		}
	}
}