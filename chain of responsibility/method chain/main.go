package main

import "fmt"

type Modifier interface {
	Add(m Modifier)
	Apply()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Apply() {
	if c.next != nil {
		c.next.Apply()
	}
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Apply() {
	fmt.Printf("Doubling %s attack\n", d.creature.Name)
	d.creature.Attack *= 2
	d.CreatureModifier.Apply()
}

func NewDoubleAttackModifier(creature *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{creature: creature},
	}
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func (i *IncreaseDefenseModifier) Apply() {
	if i.creature.Attack <= 2 {
		fmt.Printf("Increasing %s defense\n", i.creature.Name)
		i.creature.Defense++
	}
	i.CreatureModifier.Apply()
}

func NewIncreaseDefenseModifier(creature *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{
		CreatureModifier: CreatureModifier{creature: creature},
	}
}

type NoBonusesModifier struct {
	CreatureModifier
}

func (n *NoBonusesModifier) Apply() {
	// empty
}

func NewNoBonusesModifier(creature *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{
		CreatureModifier: CreatureModifier{creature: creature},
	}
}

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  attack,
		Defense: defense,
	}
}

func main() {
	murloc := NewCreature("Murloc", 1, 1)
	fmt.Println(murloc.String())

	root := NewCreatureModifier(murloc)
	root.Add(NewNoBonusesModifier(murloc))
	root.Add(NewDoubleAttackModifier(murloc))
	root.Add(NewIncreaseDefenseModifier(murloc))
	root.Add(NewDoubleAttackModifier(murloc))
	root.Apply()

	fmt.Println(murloc.String())
}
