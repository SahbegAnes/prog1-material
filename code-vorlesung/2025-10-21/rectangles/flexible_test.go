package rectangles

import "fmt"

func ExampleDrawRectangle() {
	DrawRectangle(3, 3, " ", "#")
	fmt.Println()
	DrawRectangle(3, 5, "A", "B")
	// Output:
	// ###
	// # #
	// ###
	//
	// BBBBB
	// BAAAB
	// BBBBB
}
