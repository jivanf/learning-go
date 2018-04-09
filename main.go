package main

func main() {
	d := newDeck()
	deck.print(d)
}

func newCard() string {
	return "Ace of Spades"
}

func newDeck() deck {
	deck := deck {"Ace of Diamonds", newCard()}
	deck = append(deck, "Five of Diamonds")
	return deck
}