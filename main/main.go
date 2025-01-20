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

	urlShortener := shortener.NewURLShortener(stor)

	if *shortenFlag != "" {
		shortenedURL, err := urlShortener.ShortenURL(*shortenFlag)
		if err != nil {
			log.Fatalf("Error shortening URL: %v", err)
		}
		fmt.Println("Shortened URL:", shortenedURL)
		return
	}

	if *expandFlag != "" {
		originalURL, err := urlShortener.ExpandURL(*expandFlag)
		if err != nil {
			fmt.Println("Error expanding URL or URL not found.")
			return
		}
		fmt.Println("Original URL:", originalURL)
		return
	}

	printUsage()
}

func printUsage() {
	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(1)
}
