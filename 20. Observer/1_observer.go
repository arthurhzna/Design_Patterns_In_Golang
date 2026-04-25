package main

import (
	"container/list" // list = package container/list
	"fmt"
)

type Observable struct {
	subs *list.List // --? list = package container/list
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}

func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	// let's test it!
	p.CatchACold()

	p.Unsubscribe(ds) // unsubscribe the doctor service
	p.CatchACold()    // doctor will not be called
}

// container/list.List
// Front()                         Back()
// |                                |
// v                                v
// +-------+    +-------+    +-------+
// | Value |    | Value |    | Value |
// |   A   |<-->|   B   |<-->|   C   |
// +-------+    +-------+    +-------+
// ^            ^            ^
// |____________|____________|
// 	 Prev / Next links (doubly linked)

// z := list.Front() // → node A
// z.Next() // → B
// z.Next().Next() // → C
// C.Next() // → nil
