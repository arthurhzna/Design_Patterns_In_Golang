package main

import "fmt"

type Handler interface {
	Handle(amount int)
}

type Manager struct {
	next Handler
}

func (m *Manager) Handle(amount int) {
	if amount <= 1000 { //<---- guard clause
		fmt.Println("Manager approved")
		return
	}
	m.next.Handle(amount) //<---- pass the request to the next handler, CEO will handle the request
}

type CEO struct{}

func (c *CEO) Handle(amount int) {
	fmt.Println("CEO approved")
}

func main() {
	ceo := &CEO{}
	manager := &Manager{next: ceo}

	manager.Handle(500)  // Manager approved
	manager.Handle(5000) // CEO approved
}
