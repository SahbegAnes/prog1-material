package gameboard

import "fmt"

func ExampleNewBoard() {
	fmt.Println(NewBoard())

	// Output:
	// [[     ] [     ] [     ]]
}

func ExamplePrintBoard() {
	PrintBoard(NewBoard())
	//PrintBoard([][]string{{"1", "2"}, {"3", "4"}})

	// Output:
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
}
