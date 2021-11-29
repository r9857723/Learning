package	main

import "fmt"

// 迷宮

func setWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 { //可走的路
			myMap[i][j] = 2
			if setWay(myMap, i + 1, j) { //向下
				return true
			} else if setWay(myMap, i, j + 1) { // 向右
				return true
			} else if setWay(myMap, i - 1, j) { // 向上
				return true
			} else if setWay(myMap, i, j - 1) { // 向左
				return true
			} else  {
				myMap[i][j] = 3
				return false
			}
		} else { //是牆
			return false
		}
	}
}

func main() {
	// 1 牆
	// 0 沒走過的點
	// 2 走過的點
	// 3 走過但不通
	var myMap [8][7]int
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	myMap[3][1] = 1
	myMap[3][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	//testing
	setWay(&myMap, 1, 1)
	fmt.Println("way ok")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}