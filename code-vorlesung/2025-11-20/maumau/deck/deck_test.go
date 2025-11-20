package deck

import "fmt"

func ExampleNew32() {
	deck := New32()
	fmt.Println("Anzahl Karten im Deck:", len(deck.Cards))

	// Output:
	// Anzahl Karten im Deck: 32
}
