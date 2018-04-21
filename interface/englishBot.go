package main

import "log"

type englishBot struct {
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func printGreeting(eb englishBot) {
	log.Println(eb.getGreeting())
}