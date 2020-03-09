package main

import "fmt"

/**
当一个数组中大部分元素为0，或者为同一个值的数组时，可以使用稀疏数组来保存数组

处理方法：
1） 记录数组一共有几行几列，有多少不同的值
2) 把具有不同值的行列以及值记录在一个小规模的数组中，从而缩小程序的规模

row col var
*/

type Sa struct {
	Row int
	Col int
	Var int
}

func main() {
	// 1. 创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子 *
	chessMap[2][3] = 2 // 白 o

	// 输出
	for _, v := range chessMap {
		for _, v2 := range v {
			if v2 == 1 {
				fmt.Printf("*\t")
			} else {
				if v2 == 2 {
					fmt.Printf("o\t")
				} else {
					fmt.Printf("&\t")
				}
			}
		}
		fmt.Println()
	}

	// 遍历
	var sliChe []*Sa
	//
	m := len(chessMap)
	if m < 0 {
		return
	}
	n := len(chessMap[0])
	sliChe = append(sliChe, &Sa{m, n, 0})

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				sliChe = append(sliChe, &Sa{i, j, v2})
			}
		}
	}

	for _, v := range sliChe {
		fmt.Printf("%v\n", *v)
	}

}
