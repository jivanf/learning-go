package main

import "log"

func main() {
	hand, deck := deal(5, newDeck())

	log.Println("* Cards in hand: ")
	log.Println(hand.toString())

	log.Println("* Cards left in deck: ")
	log.Println(deck.toString())
}
