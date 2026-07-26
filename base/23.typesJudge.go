package main

import (
	"fmt"
	"reflect"
)

func ThisIsFunction() {
	fmt.Println("This is a function")
}

type ThisIsType struct {
	Name string
	Age  int
}

type ThisIsMethod func()

func (t *ThisIsMethod) BaseMethod() {
	fmt.Println("This is a method")
}

type ThisIsInterface interface {
}

type ThisIsDefinedType int

func (t ThisIsDefinedType) BaseMethod() {
	fmt.Printf("This is a ThisIsDefinedType method %d", t)
}

func main() {
	a := ThisIsFunction
	b := ThisIsType{Name: "John", Age: 30}
	c := ThisIsMethod(nil)
	d := ThisIsInterface(0)
	e := ThisIsDefinedType(1)

	e.BaseMethod()

	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(ThisIsFunction))
	fmt.Println(reflect.TypeOf(b) == reflect.TypeOf(ThisIsType{Name: "John", Age: 30}))
	fmt.Println(reflect.TypeOf(c) == reflect.TypeOf(ThisIsMethod(nil)))
	fmt.Println(reflect.TypeOf(d) == reflect.TypeOf(ThisIsInterface(0)))
	fmt.Println(reflect.TypeOf(e) == reflect.TypeOf(ThisIsDefinedType(0)))
}
