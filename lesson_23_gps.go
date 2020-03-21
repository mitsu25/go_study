The Go Playground   Imports 
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	name      string
	lat, long float64
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1f°, %.1f°)", l.name, l.lat, l.long)
}

type gps struct {
	world       world
	current     location
	destination location
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destination.description())
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location{"Elysium Plantia", 4.5, 135.9}

	gps := gps{
		world:       mars,
		current:     bradbury,
		destination: elysium,
	}

	curiosity := rover{
		gps: gps,
	}
	fmt.Println(curiosity.message())
}

