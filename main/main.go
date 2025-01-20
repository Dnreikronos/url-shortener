package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Dnreikronos/url-shortener/shortener"
	"github.com/Dnreikronos/url-shortener/storage"
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

			fmt.Print("Original URL:", originalURL)
		} else {
			fmt.Println("Shortened URL not foud")
		}
		return
	}

	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(1)
}

