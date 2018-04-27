package main

import (
	"net/http"
	"log"
	"os"
	"io"
)

type logWriter struct {

}

func (logWriter ) Write(bs []byte) (int, error) {
	log.Println(string(bs))
	log.Println("Bytes that were written: ", len(bs))
	return len(bs), nil
}

func main() {
	response, err := http.Get("https://www.google.com")
	if err != nil {
		log.Println("Error: ", response)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, response.Body)
}
