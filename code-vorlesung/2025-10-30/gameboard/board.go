package gameboard

import (
	"fmt"
	"strings"
)

// Verwende eine String-Matrix für das Spielfeld.
type Board [][]string

// NewBoard liefert eine 3x3-Matrix aus Strings.
func NewBoard() Board {
	board := Board{}

	// Füge drei Zeilen zu board hinzu.
	for row := 0; row < 3; row++ {
		board = append(board, []string{" ", " ", " "})
	}

	return board
}

// PrintBoard gibt ein Spielfeld menschenlesbar auf der Konsole aus.
func PrintBoard(b Board) {
	div := strings.Repeat("+---", len(b[0])) + "+"
	for row := 0; row < len(b); row++ {
		fmt.Println(div)
		for col := 0; col < len(b[row]); col++ {
			fmt.Printf("| %v ", b[row][col])
		}
		fmt.Println("|")
	}
	fmt.Println(div)
}
