package main

import "log"

func main() {
	sB := spanishBot{}
	eB := englishBot{}
	printGreeting(sB)
	printGreeting(eB)
}

func printGreeting(b bot) {
	log.Println(b.getGreeting())
}
