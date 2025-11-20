package maumau

import "fmt"

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
	return fmt.Sprintf("┌─────────┐\n│%2s       │\n│         │\n│    %s    │\n│         │\n│       %2s│\n└─────────┘",
		c.Rank.String()[:2], c.Suit.String()[:1], c.Rank.String()[:2])
}
