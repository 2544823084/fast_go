package main

import (
	"fmt"
)

func main() {
	var a, b float32
	fmt.Println("请输入数字a")
	fmt.Scanln(&a)
	fmt.Println("请输入数字b")
	fmt.Scanln(&b)
	var option string
	fmt.Println("请输入操作符")
	fmt.Scanln(&option)
	switch option {
		case "+":
			fmt.Println("%d",a + b)
		case "-":
			fmt.Println("%d",a - b)
		case "*":
			fmt.Println("%d",a * b)
		case "/":
			if (b == 0) {
				fmt.Println("除数不能为0")
			} else {
				fmt.Println("%d", a / b)
			}
		default:
			fmt.Println("无效的操作符")
	}
}