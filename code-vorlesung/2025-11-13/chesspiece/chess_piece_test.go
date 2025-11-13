package chesspiece

import "fmt"

func ExampleChessPiece_MoveAllowed() {
	l1 := ChessPiece{
		pieceType: BISHOP,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

	fmt.Println(l1.MoveAllowed(3, 4))
	fmt.Println(l1.MoveAllowed(0, 0))

	// Output:
	// false
	// true
}
