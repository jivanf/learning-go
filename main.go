package main

import "log"

func main() {
	hand, deck := deal(5, newDeck())

	log.Println("* Cards in hand: ")
	hand.print()

	log.Println("* Cards left in deck: ")
	deck.print()
}
