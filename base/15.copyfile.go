//go:build ignore

package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	sourceFile, err := os.Open("../files/read.json")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create("../files/read_copy.json")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer destinationFile.Close()
	io.Copy(destinationFile, sourceFile)
	fmt.Println("文件复制成功")
}