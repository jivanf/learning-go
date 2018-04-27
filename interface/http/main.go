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
	responseBody := make([]byte, 9999)
	response.Body.Read(responseBody)
	log.Println(string(responseBody))
}
