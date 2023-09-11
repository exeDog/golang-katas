package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	webSites := []string{

		"http://www.google.com",
		"http://www.facebook.com",
		"https://www.twitter.com",
		"https://www.youtube.com",
		"https://www.instagram.com",
	}

	c := make(chan string)

	for _, link := range webSites {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}
