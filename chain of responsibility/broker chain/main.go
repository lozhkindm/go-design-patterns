package main

import (
	"fmt"
	"sync"
)

// chain of responsibility, mediator, observer, CQS

const (
	Attack Argument = iota
	Defense
)

type Observer interface {
	Handle(query *Query)
}

type Observable interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Fire(query *Query)
}

type Argument int

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Creature struct {
	game            *Game
	Name            string
	attack, defense int
}

func (c *Creature) GetAttack() int {
	query := Query{
		CreatureName: c.Name,
		WhatToQuery:  Attack,
		Value:        c.attack,
	}
	c.game.Fire(&query)
	return query.Value
}

func (c *Creature) GetDefense() int {
	query := Query{
		CreatureName: c.Name,
		WhatToQuery:  Defense,
		Value:        c.defense,
	}
	c.game.Fire(&query)
	return query.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.GetAttack(), c.GetDefense())
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{
		game:    game,
		Name:    name,
		attack:  attack,
		defense: defense,
	}
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Handle(query *Query) {
	if query.CreatureName == d.creature.Name && query.WhatToQuery == Attack {
		query.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func NewDoubleAttackModifier(game *Game, creature *Creature) *DoubleAttackModifier {
	mod := &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{
			game:     game,
			creature: creature,
		},
	}
	game.Subscribe(mod)
	return mod
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(observer Observer) {
	g.observers.Store(observer, struct{}{})
}

func (g *Game) Unsubscribe(observer Observer) {
	g.observers.Delete(observer)
}

func (g *Game) Fire(query *Query) {
	g.observers.Range(func(observer, _ interface{}) bool {
		if observer == nil {
			return false
		}
		observer.(Observer).Handle(query)
		return true
	})
}

func main() {
	game := &Game{observers: sync.Map{}}
	murloc := NewCreature(game, "Murloc", 2, 3)
	fmt.Println(murloc.String())

	doubleAttack := NewDoubleAttackModifier(game, murloc)
	fmt.Println(murloc.String())

	_ = doubleAttack.Close()
	fmt.Println(murloc.String())
}
