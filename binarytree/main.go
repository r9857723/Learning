package main

import (
	"fmt"
)

type Person struct {
	No int
	Name string
	Left *Person
	Right *Person
}

// PreOrder 前序遍歷
// 由上至下 由左至右
func PreOrder(root *Person) {
	if root != nil {
		fmt.Printf("no=%d name=%s\n", root.No, root.Name)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

// InfixOrder 中序遍歷
// 先輸出root的左 在輸出root 在輸出右
func InfixOrder(root *Person) {
	if root != nil {
		InfixOrder(root.Left)
		fmt.Printf("no=%d name=%s\n", root.No, root.Name)
		InfixOrder(root.Right)
	}
}

// PostOrder 後序遍歷
// 先輸出root的左 在輸出root的右 在輸root
func PostOrder(root *Person) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Printf("no=%d name=%s\n", root.No, root.Name)
	}
}

func main() {
	root := &Person{ No: 1, Name: "Tom" }

	left1 := &Person{ No: 2, Name: "Roman" }

	left2 := &Person{ No: 21, Name: "Q" }
	right3 := &Person{ No: 22, Name: "W" }

	left1.Left = left2
	left1.Right = right3

	right1 := &Person{ No: 3, Name: "Yuki" }

	root.Left = left1
	root.Right = right1

	right2 := &Person{ No: 31, Name: "Michael" }

	right1.Right = right2

	//PreOrder(root)
	//InfixOrder(root)
	PostOrder(root)
}