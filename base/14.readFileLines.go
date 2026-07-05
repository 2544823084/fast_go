package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	file, err := os.Open("../files/read.json")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	fmt.Println("文件行数：", count)
}