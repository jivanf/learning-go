package main

import (
	"log"
	"strings"
	"io/ioutil"
)
type deck []string

func (d deck) print() {
	for _, card := range d {
		log.Println(card)
	}
}

func newDeck() deck {
	d := deck{}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			d = append(d, cardValue + " of " + cardSuit)
		}
	}
	return d
}

func deal(cards int, d deck) (deck, deck) {
	return d[:cards], d[cards:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}