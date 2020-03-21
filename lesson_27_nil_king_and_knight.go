package main

import (
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v gets a %v\n", c.name, i.name)
	c.leftHand = i
}

func (from *character) give(to *character) {
	if from == nil || to == nil || from.leftHand == nil {
		return
	}
	to.leftHand = from.leftHand
	fmt.Printf("%v gives %v to %v\n", from.name, from.leftHand.name, to.name)
	from.leftHand = nil
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v has nothing.", c.name)
	}
	return fmt.Sprintf("%v has a %v", c.name, c.leftHand.name)
}

func main() {
	king := &character{name: "King"}
	knight := &character{name: "Knight"}
	item := &item{name: "sword"}

	king.pickup(item)

	king.give(knight)

	fmt.Println(knight)
	fmt.Println(king)
}

