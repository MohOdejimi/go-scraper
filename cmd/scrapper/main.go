package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"goWebScrapper/internal/scrapper"
)

func main() {
	urlFlag := flag.String("url", "", "The URL to scrape content from")
	outputFlag := flag.String("output", "output", "The filename (without extension) to save the content")
	formatFlag := flag.String("format", "txt", "The format to save the content")
	portFlag := flag.String("port", "", "The port to serve the file for download")

	flag.Parse()

	if *urlFlag == "" {
		log.Fatal("Error: Please provide a URL using the -url flag")
		flag.Usage()
		os.Exit(1)
	}

	if *portFlag == "" {
		log.Fatal("Error: Please provide a port using the -port flag")
		flag.Usage()
		os.Exit(1)
	}
 
	doc, err := scrapper.Fetch(*urlFlag)
	if err != nil {
		log.Fatalf("Error fetching URL: %s", err)
	}

 
	parsedData, err := scrapper.ParseHTML(doc)
	if err != nil {
		log.Fatalf("Error parsing HTML: %s", err)
	}
 
	filePath := scrapper.DetermineFileFormat(*formatFlag, *outputFlag, parsedData)
	fmt.Printf("Scraped data saved to: %s\n", filePath)
 
	go func() {
		http.Handle("/", http.FileServer(http.Dir(".")))
		fmt.Printf("Download your file at: http://localhost:%s/%s\n", *portFlag, filePath)
		log.Fatal(http.ListenAndServe(":"+*portFlag, nil))
	}()

	select {}  
}
