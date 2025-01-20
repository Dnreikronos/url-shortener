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
	shortenFlag := flag.String("shorten", "", "URL to shorten")
	expandFlag := flag.String("expand", "", "Shortened URL to expand")
	flag.Parse()


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

