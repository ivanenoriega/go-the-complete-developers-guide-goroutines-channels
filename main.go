package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://udemy.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is ok")
	c <- link
}
