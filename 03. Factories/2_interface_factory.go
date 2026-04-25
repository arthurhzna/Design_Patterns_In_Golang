package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
}

func NewPerson(name string, age int) Person { //factory function
	return &person{name, age}
}

// Returns a pointer to the Person struct.
// NewPerson returns a struct that has the SayHello method.

func main() {
	p := NewPerson("John", 20)
	p.SayHello()
}
