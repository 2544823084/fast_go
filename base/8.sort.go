//go:build ignore

package main

import (
	"fmt"
)

func sort(arr []int) []int {
	for i:=0; i<len(arr); i++ {
		for j:=0; j<len(arr)-i-1; j++ {
			if (arr[j] > arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	var len int
	fmt.Scanln(&len)
	arr := make([]int, len)
	for i:=0; i<len; i++ {
		fmt.Scanln(&arr[i])
	}
	var newArr = sort(arr)
	fmt.Println(newArr)
}