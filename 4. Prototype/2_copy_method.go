package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address { // new struct, why ? because we want to return a new address struct, if jane change address, john's address should not change
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country}
}

// Copies John's values, not the pointer.
// Because a is *Address, using a.StreetAddress copies the value of john.StreetAddress,
// not a pointer to john.StreetAddress.

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p // copies Name
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends) // copy the friends slice
	return &q
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
