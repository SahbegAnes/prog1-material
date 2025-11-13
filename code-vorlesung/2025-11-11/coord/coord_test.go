package coord

import (
	"fmt"
	"math"
)

type Coord struct {
	Longitude float64
	Latitude  float64
}

// Berechnet den Abstand zwischen den beiden Koordinaten c und o.
func (c Coord) Distance(o Coord) float64 {
	dx := c.Latitude - o.Latitude
	dy := c.Longitude - o.Longitude
	h := math.Sqrt(dx*dx + dy*dy)
	return h
}

func ExampleCoord_usage() {
	//a := []int{1, 2}
	a := Coord{1, 2.5}
	b := Coord{3, 4.5}

	fmt.Println(a)
	fmt.Println(a.Longitude)
	fmt.Println(a.Latitude)
	fmt.Println(a.Distance(b))

	// Output:
}
