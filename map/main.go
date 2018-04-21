package main

import "log"

func main() {
	 colors := map[string]string {
	 	"red":"#ff0000",
	 	"green":"#008000",
	 	"blue":"#0000ff",
	 }

	 printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		log.Println(k, "is", v)
	}
}
