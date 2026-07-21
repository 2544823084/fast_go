//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) // 信号通道

	go func() {
		for i := 1; i <= 10; i += 2 { // 奇数: 1,3,5,7,9
			fmt.Printf("奇数: %d\n", i)
			<-done // 等通知再继续
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 { // 偶数: 2,4,6,8,10
			fmt.Printf("偶数: %d\n", i)
			done <- struct{}{} // 通知对方可以打印了
		}
	}()

	done <- struct{}{} // 先给偶数一个启动信号
	time.Sleep(10 * time.Second)
	// 需要 WaitGroup 或 time.Sleep 等 goroutine 结束
	// （实际项目用 sync.WaitGroup）
}
