package main

import "log"

func main() {
	 colors := map[string]string {
	 	"red":"#ff0000",
	 	"green":"#008000",
	 }

	 emptyMap1 := make(map[string]string)
	 var emptyMap2 map[string]string

	 log.Println(colors)
	 log.Println("emptyMap1", emptyMap1)
	 log.Println("emptyMap2", emptyMap2)
}
