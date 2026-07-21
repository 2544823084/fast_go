//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	sum := N*(N+1)/2
	fmt.Println(sum)
}