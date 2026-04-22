package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier // singly linked list
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil { //<-- call from root until the end of the chain, root -> handler 1 -> handler 2 -> handler 3 -> nil
		c.next.Add(m) //<--- embeded, DoubleAttackModifier have add method, so it will call the add method of the next handler
	} else {
		c.next = m
	}
}

// (c *CreatureModifier) Add( --> call c.next.Add(m)
// Step 1

// Receiver c = root

// c.next.Add(m)

// ➡️ call H1.Add(m)

// Step 2

// Receiver c = H1

// c.next.Add(m)

// ➡️ call H2.Add(m)

// Step 3

// Receiver c = H2

// c.next.Add(m)

// ➡️ call H3.Add(m)

// Step 4

// Receiver c = H3

// If:

// H3.next == nil

// attach new node:

// H3.next = newNode

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(
	c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{
		creature: c}}
}

type IncreasedDefenseModifier struct {
	CreatureModifier
}

func NewIncreasedDefenseModifier(
	c *Creature) *IncreasedDefenseModifier {
	return &IncreasedDefenseModifier{CreatureModifier{
		creature: c}}
}

func (i *IncreasedDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing",
			i.creature.Name, "\b's defense")
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name,
		"attack...")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(
	c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{
		creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	// nothing here!
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	//root.Add(NewNoBonusesModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreasedDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	// 	root.next
	// 	↓
	//  DoubleAttackModifier

	//  DoubleAttackModifier.next
	// 	↓
	//  IncreasedDefenseModifier

	//  IncreasedDefenseModifier.next
	// 	↓
	//  nil

	// eventually process the entire chain
	root.Handle()
	fmt.Println(goblin.String())
}
