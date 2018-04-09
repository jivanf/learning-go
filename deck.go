package main

import "log"

type deck []string

func (d deck) print() {
	for i, card := range d {
		log.Println(i, card)
	}
}