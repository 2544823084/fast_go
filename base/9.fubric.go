package main

import (
	"fmt"
)

func fubric(n int) int {
	if (n == 0) {
		return 0
	}
	if (n == 1) {
		return 1
	}
	with := make([]int, n+1)
	with[0] = 0
	with[1] = 1
	for i:=2; i<=n; i++ {
		with[i] = with[i-1] + with[i-2]
	}
	return with[n]
}

func main() {
	var n int
	fmt.Scanln(&n)
	var result = fubric(n)
	fmt.Println(result)
}