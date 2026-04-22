package main

import "fmt"

type Handler interface {
	Handle(amount int)
}

// Handler 1
type Employee struct {
	next Handler
}

func (e *Employee) Handle(amount int) {
	if amount <= 100 {
		fmt.Println("Employee approved")
		return
	}

	fmt.Println("Employee passed request")
	e.next.Handle(amount)
}

// Handler 2
type Manager struct {
	next Handler
}

func (m *Manager) Handle(amount int) {
	if amount <= 1000 {
		fmt.Println("Manager approved")
		return
	}

	fmt.Println("Manager passed request")
	m.next.Handle(amount) // <---- this is chain of responsibility, Manager passed the request to the next handler, CEO will handle the request
}

// Handler 3
type CEO struct{}

func (c *CEO) Handle(amount int) {
	fmt.Println("CEO approved")
}

func main() {
	ceo := &CEO{}
	manager := &Manager{next: ceo}
	employee := &Employee{next: manager}

	employee.Handle(50)
	fmt.Println("-----")

	employee.Handle(500)
	fmt.Println("-----")

	employee.Handle(5000)
}
