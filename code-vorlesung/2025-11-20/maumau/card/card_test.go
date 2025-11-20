package card

import "fmt"

func ExampleCard_String() {
	card := NewCard(Ace, Spades)
	fmt.Println(card.String())

	// Output:
	// Ace of Spades
}

func ExampleCard_Picture() {
	card1 := NewCard(King, Hearts)
	fmt.Println(card1.Picture())
	fmt.Println()

	card2 := NewCard(Seven, Diamonds)
	fmt.Println(card2.Picture())
	fmt.Println()

	card3 := NewCard(Ten, Clubs)
	fmt.Println(card3.Picture())

	// Output:
	// ┌─────────┐
	// │K        │
	// │         │
	// │    ♥    │
	// │         │
	// │        K│
	// └─────────┘
	//
	// ┌─────────┐
	// │7        │
	// │         │
	// │    ♦    │
	// │         │
	// │        7│
	// └─────────┘
	//
	// ┌─────────┐
	// │10       │
	// │         │
	// │    ♣    │
	// │         │
	// │       10│
	// └─────────┘
}
