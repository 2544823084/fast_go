package main

import (
	"fmt"
	"strings"
)

func reduceSlice(arr []string) []string {
	seen := make(map[string]struct{})
	for _, item := range arr {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
	}
	stringArr := make([]string, 0, len(seen))
	for item := range seen {
		stringArr = append(stringArr, item)
	}
	str := strings.Join(stringArr, " ")
	return str
}

func main() {
	var arr []int
	fmt.Scanln(&arr)
	var newArr = reduceSlice(arr)
	fmt.Println(newArr)
}