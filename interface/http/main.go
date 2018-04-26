package main

import (
	"net/http"
	"log"
	"os"
)

func main() {
	response, err := http.Get("https://www.google.com")
	if err != nil {
		log.Println("Error: ", response)
		os.Exit(1)
	}
	log.Println(response.Body)
}
