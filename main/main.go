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

	stor, err := storage.NewGormStorage()
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

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

