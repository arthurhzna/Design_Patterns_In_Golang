package main

import (
	"fmt"
	"sync"
)

// cqs, mediator, cor

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(*Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
	//                   ↑↑↑ empty anon struct
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool { //--> for loop map from game.observers  : sync.Map
		if key == nil {
			return false
		}
		key.(Observer).Handle(q) //--> call the handle method of the observer
		return true
	})
}

type Creature struct { // define character with attack and defense
	game            *Game
	Name            string
	attack, defense int // ← private!
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

// Every call to Attack() or Defense() creates a new Query and fires it through the game.
// Remember: this is not a pointer.
// SO AFTER, UNSUBSCRIBE THE MODIFIER
// m := NewDoubleAttackModifier(game, goblin)
// fmt.Println(goblin.String())
// m.Close()
// fmt.Println(goblin.String())
// --> the double attack modifier is removed, goblin will have original attack and defense

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack(), c.Defense())
}

// data common to all modifiers
type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c *CreatureModifier) Handle(*Query) {
	// nothing here!
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name &&
		q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func main() {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2) // define character with attack and defense
	fmt.Println(goblin.String())
	fmt.Println("-----")

	{
		m := NewDoubleAttackModifier(game, goblin) // add double attack modifier to the character, game name is "game", character name is "Strong Goblin"
		fmt.Println(goblin.String())               // goblin will have double attack modifier
		m.Close()                                  // remove the modifier
	} // --> the double attack modifier is removed, goblin will have original attack and defense

	// CANNOT USE m, after } block, m is out of scope
	fmt.Println(goblin.String())
}
