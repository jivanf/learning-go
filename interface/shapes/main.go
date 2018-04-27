package main

import "log"

func main() {
	sq := square{10}
	t := triangle{10, 10}
	printArea(sq)
	printArea(t)
}

func printArea(s shape) {
	log.Println(s.getArea())
}
