package card

import (
	"fmt"
	"strings"
)

// Card repräsentiert eine Spielkarte mit einem Rang und einer Farbe.
type Card struct {
	Rank
	Suit
}

// NewCard erstellt eine neue Karte mit einem Rang und einer Farbe.
func NewCard(rank Rank, suit Suit) Card {
	return Card{
		Rank: rank,
		Suit: suit,
	}
}

// String gibt eine lesbare Darstellung der Karte zurück.
func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.Rank, c.Suit)
}

// Picture gibt eine ASCII-Art-Darstellung der Karte zurück.
func (c Card) Picture() string {
	lines := []string{
		"┌─────────┐",
		fmt.Sprintf("│%-2s       │", c.Rank.String()),
		"│         │",
		fmt.Sprintf("│    %s    │", c.Suit.Symbol()),
		"│         │",
		fmt.Sprintf("│       %2s│", c.Rank.String()),
		"└─────────┘",
	}
	return strings.Join(lines, "\n")
}

// Matches erwartet eine zweite Karte und liefert true,
// falls diese zur aktuellen Karte passt.
func (c Card) Matches(o Card) bool {
	// TODO
	return true
}
