package main

import (
	"fmt"
	"math/rand"
	"time"
)

type creature interface {
	eat() string
	move() string
}

type elephant struct {
	name string
}

func (e elephant) String() string {
	return e.name
}

func (e elephant) move() string {
	switch rand.Intn(3) {
	case 0:
		return "doshi doshi"
	case 1:
		return "dosshing! dosshing!"
	default:
		return "........."

	}
}

func (e elephant) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "apple"
	case 1:
		return "cabbage"
	default:
		return "carrot"

	}
}

type penguin struct {
	name string
}

func (p penguin) String() string {
	return p.name
}

func (p penguin) move() string {
	switch rand.Intn(3) {
	case 0:
		return "peta peta"
	case 1:
		return "pyon pyon"
	default:
		return "gyuiiiiiinnnnn!!!"

	}
}

func (p penguin) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "fish"
	case 1:
		return "big fish"
	default:
		return "super big fish!!"

	}
}

func action(c creature) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v is moving: %v\n", c, c.move())
	default:
		fmt.Printf("%v is eating a %v\n", c, c.eat())
	}
}

const sunrise, sunset int = 5, 19

func main() {
	rand.Seed(time.Now().UnixNano())

	creatures := []creature{
		elephant{name: "Dambo"},
		penguin{name: "Pingoo"},
	}

	var day, hour int = 1, 0

	for {
		fmt.Printf("Day %d %2d:00 : ", day, hour)
		if sunrise < hour && hour < sunset {
			action(creatures[rand.Intn(len(creatures))])
		} else {
			fmt.Println("The creatures are spleeping.....")
		}

		hour++
		if hour == 24 {
			hour = 0
			day++
			if day == 3 {
				break
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

