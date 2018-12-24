package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := getUrls()
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			sleep()
			checkLink(link, c)
		}(l)
	}
}

func getUrls() []string {
	return []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"https://stackoverflow.com",
	}
}

func checkLink(link string, c chan (string)) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

func sleep() {
	time.Sleep(5 * time.Second)
}
