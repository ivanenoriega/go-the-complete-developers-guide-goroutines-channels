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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is ok")
}
