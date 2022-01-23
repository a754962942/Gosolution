package main

import "fmt"

/*

二位数组判断
数组按行递增，按列递增
要求：输入一个二维数组并输入一个数字，判断数组内是否存在该数字，
存在返回true ，否则返回false
*/

func Find(arr [][]int, target int) bool {
	if len(arr) > 0 && len(arr[0]) > 0 {
		//定义行和列
		r := 0
		c := len(arr[0]) - 1
		for r < len(arr) && c >= 0 {
			if target < arr[r][c] {
				c--
			} else if target > arr[r][c] {
				r++
			} else if target == arr[r][c] {
				return true
			}
		}
	}
	return false
}

func main() {
	var arr = [][]int{
		{1, 2, 8, 9, 11},
		{2, 4, 9, 12, 14},
		{4, 7, 10, 13, 17},
		{6, 8, 11, 15, 20},
	}
	b := Find(arr, 3)
	fmt.Println(b)
	// fmt.Println(arr)
	// fmt.Println(len(arr))
	// fmt.Println(len(arr[0]))

}
