package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/Dnreikronos/url-shortener/shortener"
)

var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)

func main() {

	shorten := flag.String("Shorten", "", "Url to shorten")
	expand := flag.String("Expand", "", "Shortened URL to Expand")

	flag.Parse()

	if *shorten != "" {
		shortenedURL := shortener.ShortenURL(*shorten)
		fmt.Println("Shortened URL:", shortenedURL)
		return
	}

	if *expand != "" {
		originalURL := shortener.ExpandURL(*expand)
		if originalURL != "" {

			fmt.Printf("Original URL:", originalURL)
		} else {
			fmt.Println("Shortened URL not foud")
		}
		return
	}

	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(1)
}

