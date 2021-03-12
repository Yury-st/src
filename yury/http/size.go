package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {

	pages := make(chan Page)
	urls := []string{"https://vk.com/adkazy", "https://golang.org", "https://google.com"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i, _ := range urls {
		page := <-pages
		fmt.Printf("%v %v %v \n", i, page.URL, page.Size)
	}
}

func responseSize(url string, pages chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	pages <- Page{url, len(body)}
}
