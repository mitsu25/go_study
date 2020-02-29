package main

import (
	"fmt"
)

var planets Planets

type Planets []string

func (planets Planets) terraform() {
	for i, planet := range planets {
		if planet == "Mars" || planet == "Uranus" || planet == "Neptune" {
			planets[i] = "New" + planet
		}
	}
}

func main() {
	planets = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	planets.terraform()
	fmt.Print(planets)
}

