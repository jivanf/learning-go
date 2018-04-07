package main

import (
	"log"
)

func main() {
	deck := newDeck()
	for i, card := range deck {
		log.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}

func newDeck() []string {
	deck := []string {"Ace of Diamonds", newCard()}
	deck = append(deck, "Five of Diamonds")
	return deck
}