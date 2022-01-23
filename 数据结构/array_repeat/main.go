/*
找出数组中重复的数字
在一个长度为n的数组的所有数字都在0~n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复
也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字：
例如   如果输入长度为7的数组{2，3，1，0，2，5，3}，那么对应的输出是重复数字的2或者3
*/
package main

import (
	"fmt"
)

//下标和值对应
//时间复杂度O(n)
func indexsort(arr []int) int {
	//for循环遍历数组

	return -1
}
func main() {
	var arr = []int{2, 3, 1, 0, 2, 5, 3}
	//法1
	//O(n*log n)[2, 3, 5, 4, 3, 2, 6, 7
	//sort.Ints(arr)
	// fmt.Println(arr)

	//法2
	//时间复杂度为O(n)，找出所有
	m := make(map[int]int, len(arr))
	for _, v := range arr {
		m[v] += 1
		if m[v] > 1 {
			fmt.Println(v)
		}
	}

	//法3
	//时间复杂度O(1)
	// for i := 0; i < len(arr)-1; i++ {
	// 	//如果条件不满足进入循环
	// 	for arr[i] != i {
	// 		//如果arr[i] ==arr[arr[i]]返回arr[i]
	// 		//否则冒泡排序，再次循环
	// 		if arr[i] == arr[arr[i]] {
	// 			fmt.Println(arr[i])
	// 			return
	// 		} else {
	// 			//冒泡排序原理
	// 			temp := arr[i]
	// 			arr[i] = arr[temp]
	// 			arr[temp] = temp
	// 		}
	// 	}
	// }

}
