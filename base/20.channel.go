//go:build ignore

package main

import "fmt"

func sumPart(nums []int, ch chan int) {
	total := 0
	for _, n := range nums {
		total += n
	}
	ch <- total // 将结果发送到 channel
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 分成两半
	mid := len(nums) / 2
	left := nums[:mid]  // [1,2,3,4,5]
	right := nums[mid:] // [6,7,8,9,10]

	ch := make(chan int) // 用一个 channel 收集两个 goroutine 的结果

	go sumPart(left, ch)
	go sumPart(right, ch)

	// 读取两个结果并汇总
	sum1 := <-ch
	sum2 := <-ch

	fmt.Printf("左半部分和: %d\n", sum1)
	fmt.Printf("右半部分和: %d\n", sum2)
	fmt.Printf("总和: %d\n", sum1+sum2)
}