package main

import (
	"log"
	"strings"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
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
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	deck := strings.Split(string(bs), ",")
	return deck
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(d); i++ {
		r := rand.Intn(len(d) - 1)
		d[i], d[r] = d[r], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}