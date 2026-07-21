//go:build ignore

package main

import (
	"fmt"
)

func countWord(str string) map[string]int {
	wordMap := make(map[string]int)
	words := []rune(str)
	for _, word := range words {
		wordMap[string(word)]++
	}
	return wordMap
}

func main() {
	var str string
	fmt.Scanln(&str)
	var wordMap = countWord(str)
	fmt.Println(wordMap)
}