package main

import (
	"os"
	"log"
	"io"
)

func main() {
	file, err := os.Open(os.Args[1])
	bs := make([]byte, 99999)
	if err == nil {
		io.Copy(file, file)
	}
	log.Println(string(bs))

}
