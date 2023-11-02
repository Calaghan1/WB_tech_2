// Утилита wget

// Реализовать утилиту wget с возможностью скачивать сайты целиком.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"golang.org/x/net/html"
)
func CheckErorr(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		return true
	}
}

func CollectAllLinks(url string) {
	all_lincs := make(map[string]bool)
	fmt.Println(url)
	resp, err := http.Get(url)
	if !CheckErorr(err) {
		os.Exit(0)
	}
	doc, err := html.Parse(resp.Body)
	if !CheckErorr(err) {
		os.Exit(0)
	}
	if doc.Type == html.ElementNode && doc.Data == "a" {
		
	}
	fmt.Println(resp.Body)
}

func main() {
	url := os.Args[0]
	CollectAllLinks(url)
}
