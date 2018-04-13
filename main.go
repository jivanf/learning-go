package main

func main() {
	hand, _ := deal(5, newDeck())
	hand.shuffle()
	hand.print()
}
