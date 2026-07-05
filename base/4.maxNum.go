package main

import (
	"fmt"
)
func maxNum(arr []int) int {
	var max = arr[0]
	for i:=1; i<len(arr); i++ {
		if (max < arr[i]) {
			max = arr[i]
		}
	}
	return max
}
func main()  {
	var len int
	fmt.Println("请输入数组长度：")
	fmt.Scan(&len)
	arr := make([]int, len)
	for i:=0 ;i < len; i++ {
		fmt.Scan(&arr[i])
	}
	max := maxNum(arr)
	fmt.Println(max)
}
