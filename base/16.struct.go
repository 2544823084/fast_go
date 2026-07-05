package main

import (
	"fmt"
)

type Rectangle struct {
	Width float32
	Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Width + r.Height)
}

func main() {
	var width, height float32
	fmt.Println("请输入矩形的宽")
	fmt.Scanln(&width)
	fmt.Println("请输入矩形的高")
	fmt.Scanln(&height)
	rectangle := Rectangle{Width: width, Height: height}
	fmt.Println("矩形的面积是：", rectangle.Area())
	fmt.Println("矩形的周长是：", rectangle.Perimeter())
}