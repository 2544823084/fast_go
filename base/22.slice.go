//go:build ignore

package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
  slice := arr[1:]
  slice[0] = 0
  fmt.Printf("array: %v\n", arr)
  fmt.Printf("slice: %v\n", slice)
}