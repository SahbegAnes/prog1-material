package rectangles

import "fmt"

func ExampleDrawEmptyRectangle() {
	DrawEmptyRectangle(3, 3)
	fmt.Println()
	DrawEmptyRectangle(3, 6)
	// Output:
	// ###
	// # #
	// ###
	//
	// ######
	// #    #
	// ######

}
