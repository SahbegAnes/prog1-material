package card

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// String gibt eine lesbare Darstellung der Farbe zurück.
func (s Suit) String() string {
	return s.Name()
}

// Name gibt einen menschlich lesbaren Namen der Farbe zurück.
func (s Suit) Name() string {
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

// Symbol gibt das Symbol für die Farbe zurück.
func (s Suit) Symbol() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	default:
		return "?"
	}
}
