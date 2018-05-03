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

	c1 := make(chan string)
	c2 := make(chan string)

	for {
		for k, v := range websites {
			go checkLink(k, v, c1, c2)
			log.Println(<-c1, <-c2)

		}
	}

}

func checkLink(name string, link string, c1 chan string, c2 chan string) {
	resp, err := http.Get(link)
	c1 <- name
	c2 <- resp.Status
	if err != nil {
		log.Println(name,"might be down!")
		return
	}


}
