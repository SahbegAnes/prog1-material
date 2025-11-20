package deck

import (
	"fmt"
	"prog1-material/code-vorlesung/2025-11-20/maumau/card"
)

func ExampleNew32() {
	deck := New32()
	fmt.Println("Anzahl Karten im Deck:", len(deck.Cards))

	// Output:
	// Anzahl Karten im Deck: 32
}

func ExampleDeck_Draw() {
	deck := New32()
	card, err := deck.Draw()
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Println("Gezogene Karte:", card)
	fmt.Println("Karten übrig im Deck:", len(deck.Cards))

	// Output:
	// Gezogene Karte: A of Spades
	// Karten übrig im Deck: 31
}

func ExampleDeck_Add() {
	deck := New32()

	card := card.NewCard(card.Ten, card.Hearts)
	deck.Add(card)

	fmt.Println("Karten nach Hinzufügen im Deck:", len(deck.Cards))
	c2, _ := deck.Draw()
	fmt.Println("Gezogene Karte nach Hinzufügen:", c2)

	// Output:
	// Karten nach Hinzufügen im Deck: 33
	// Gezogene Karte nach Hinzufügen: 10 of Hearts
}
