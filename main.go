package main

import (
	"log"
)

func main() {
	card := newCard()
	log.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
