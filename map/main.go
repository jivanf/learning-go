package main

import "log"

func main() {
	 colors := map[string]string {
	 	"red":"#ff0000",
	 	"green":"#008000",
	 }

	 colors["blue"] = "#0000ff"

	 log.Println(colors)
}
