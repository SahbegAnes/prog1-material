package rectangles

import "fmt"

func ExampleDrawSolidRectangle() {
	DrawSolidRectangle(3, 3)
	fmt.Println()
	DrawSolidRectangle(3, 6)
	// Output:
	// ###
	// ###
	// ###
	//
	// ######
	// ######
	// ######

}
