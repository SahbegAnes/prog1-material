package deck

import (
	"fmt"
	"math/rand"
	"prog1-material/code-vorlesung/2025-11-20/maumau/card"
)

// Deck repräsentiert ein Kartendeck.
type Deck struct {
	Cards []card.Card
}

// New32 erstellt ein neues Kartendeck mit 32 Karten.
func New32() Deck {
	var cards []card.Card
	for suit := card.Clubs; suit <= card.Spades; suit++ {
		for rank := card.Seven; rank <= card.Ace; rank++ {
			cards = append(cards, card.NewCard(rank, suit))
		}
	}
	return Deck{Cards: cards}
}

// Shuffle mischt die Karten im Deck.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Draw entfernt die oberste Karte vom Deck und gibt sie zurück.
func (d *Deck) Draw() (card.Card, error) {
	if len(d.Cards) == 0 {
		return card.Card{}, fmt.Errorf("kein Karte im Deck")
	}
	last := len(d.Cards) - 1
	topCard := d.Cards[last]
	d.Cards = d.Cards[:last]
	return topCard, nil
}

// Add fügt eine Karte oben zum Deck hinzu.
func (d *Deck) Add(card card.Card) {
	d.Cards = append(d.Cards, card)
}

// Top liefert die oberste Karte, ohne sie zu entfernen.
// TODO
