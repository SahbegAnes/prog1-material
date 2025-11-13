package const_und_enums

import "fmt"

// 0: Norden
// 1: Osten
// 2: Süden
// 3: Westen
type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

// const (
// 	NORTH Direction = 0
// 	EAST  Direction = 1
// 	SOUTH Direction = 2
// 	WEST  Direction = 3
// )

type Coord struct {
	Longitude float64
	Latitude  float64
}

func (c *Coord) Step(d Direction) {
	switch d {
	case NORTH:
		c.Latitude += 1
	case EAST:
		c.Longitude += 1
	case SOUTH:
		c.Latitude -= 1
	case WEST:
		c.Longitude -= 1
	}

	/*
		if d == 0 {
			c.Latitude += 1
		} else if d == 1 {
			c.Longitude += 1
		} else if d == 2 {
			c.Latitude -= 1

		} else {
			c.Longitude -= 1
		}
	*/
}

func ExampleCoord_usage() {
	c1 := Coord{1.5, 3}
	c2 := Coord{Latitude: 4, Longitude: 5.5}

	fmt.Println(c1)
	fmt.Println(c2)

	c1.Step(NORTH)
	c1.Step(EAST)

	fmt.Println(c1)

	c1.Step(SOUTH)
	c1.Step(WEST)

	fmt.Println(c1)

	// Output:
	// {1.5 3}
	// {5.5 4}
	// {2.5 4}
	// {1.5 3}
}

/* Bessere Version von Step mit Fehlerbehandlung:
func (c *Coord) Step(d Direction) error {
	switch d {
	case 0:
		c.Latitude += 1
	case 1:
		c.Longitude += 1
	case 2:
		c.Latitude -= 1
	case 3:
		c.Longitude -= 1
	default:
		return fmt.Errorf("Geht nicht, du ...")
	}
	return nil
}
*/
