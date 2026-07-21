//go:build ignore

package main

import "fmt"

type Animal interface {
	say() string
}

type Dog struct {
	name string
}

func (d Dog) say() string {
	return d.name + " is running"
}

func main() {
	var animal Animal
	animal = Dog{name: "dog"}
	fmt.Println(animal.say())
}