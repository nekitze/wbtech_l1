package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) SayHello() {
	fmt.Printf("Hello! My name is %s.\n", h.Name)
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{Name: "Ivan"}}
	action.SayHello()
}
