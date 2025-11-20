package maumau

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// String gibt eine lesbare Darstellung der Farbe zurück.
func (s Suit) String() string {
	switch s {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return "Unknown Suit"
	}
}
