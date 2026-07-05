package main

import (
	"fmt"
)

func reverseStr(str string) string {
	rune := []rune(str)
	for i:=0; i< len(rune)/2; i++ {
		rune[i], rune[len(rune)-i-1] = rune[len(rune)-i-1], rune[i]
	}
	return string(rune)
}


func main() {
	var str string
	fmt.Scanln(&str)
	var reStr = reverseStr(str)
	fmt.Println(reStr)
}