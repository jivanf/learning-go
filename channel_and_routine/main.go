package main

import (
	"net/http"
	"log"
)

func main() {
	websites := map[string]string {
		"Google":"http://www.google.com",
		"Facebook":"http://www.facebook.com",
		"Amazon":"http://www.amazon.com",
		"Go Docs":"http://www.golang.org",
		"Stack Overflow":"http://stackoverflow.com",
	}
	for k,v := range websites {
		resp, err := http.Get(v)
		if err != nil {
			log.Println(v,"might be down!")
			continue
		}
		log.Println(k,resp.Status)

	}
}
