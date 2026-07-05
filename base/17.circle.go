package main

import (
	"fmt"
	"math"
)
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	circle := Circle{Radius: 10}
	rectangle := Rectangle{Width: 10, Height: 20}
	triangle := Triangle{Base: 10, Height: 20}
	fmt.Println("圆的面积是：", circle.Area())
	fmt.Println("矩形的面积是：", rectangle.Area())
	fmt.Println("三角形的面积是：", triangle.Area())
}