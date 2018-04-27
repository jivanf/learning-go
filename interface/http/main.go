package main

import (
	"net/http"
	"log"
	"os"
	"io"
)

func main() {
	response, err := http.Get("https://www.google.com")
	if err != nil {
		log.Println("Error: ", response)
		os.Exit(1)
	}
	io.Copy(os.Stdout, response.Body)
}
