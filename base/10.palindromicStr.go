package main

import (
	"fmt"
)

func palindromicStr (str string) bool {
	runes := []rune(str)
	for i:=0; i< len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var n string
	fmt.Scanln(&n)
	var result = palindromicStr(n)
	fmt.Println(result)
}