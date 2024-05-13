package main

import "fmt"

// First adaptee
type Human struct {
	Name string
}

func (h Human) SayHello() {
	fmt.Printf("Human: Hello, my name is %s!\n", h.Name)
}

// Second adaptee
type Alien struct{}

func (a Alien) SaySomething() {
	fmt.Println("Alien: Skibidi dop!")
}

// Target interface
type Speaker interface {
	Speak()
}

// First adapter
type HumanAdapter struct {
	*Human
}

func (h Human) Speak() {
	h.SayHello()
}

// Second adapter
type AlienAdapter struct {
	*Alien
}

func (a Alien) Speak() {
	a.SaySomething()
}

func main() {
	alien := &Alien{}
	human := &Human{Name: "Ivan"}

	speakers := [2]Speaker{HumanAdapter{human}, AlienAdapter{alien}}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}
