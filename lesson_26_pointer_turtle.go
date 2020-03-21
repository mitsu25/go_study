package main

import (
	"fmt"
	"math/rand"
)

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y++
}

func (t *turtle) down() {
	t.y--
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) left() {
	t.x--
}

func (t turtle) String() string {
	return fmt.Sprintf("location x: %d, y: %d", t.x, t.y)
}

func main() {
	t := &turtle{0, 0}

	for i := 0; i < 100; i++ {
		switch rand.Intn(4) {
		case 1:
			t.up()
		case 2:
			t.down()
		case 3:
			t.right()
		default:
			t.left()
		}
		fmt.Println(*t)
		i++
	}
	fmt.Println(*t)
}

