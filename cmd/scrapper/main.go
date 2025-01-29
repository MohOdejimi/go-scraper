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
	outputFlag := flag.String("output", "output.txt", "The file to write the scraped content to")
	portFlag := flag.String("port", "8080", "The port to serve the file for download")

	flag.Parse()

	if *urlFlag == "" {
		log.Fatal("Error: Please provide a URL using the -url flag")
		flag.Usage()
		os.Exit(1)
	}
 
	urlContent, err := scrapper.Fetch(*urlFlag)
	if err != nil {
		log.Fatalf("Error fetching URL: %s", err)
	}

	 
	err = os.WriteFile(*outputFlag, []byte(urlContent), 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	fmt.Printf("Content successfully scraped from %s and saved to %s\n", *urlFlag, *outputFlag)

	 
	go func() {
		http.Handle("/", http.FileServer(http.Dir(".")))
		fmt.Printf("Download your file at: http://localhost:%s/%s\n", *portFlag, *outputFlag)
		log.Fatal(http.ListenAndServe(":"+*portFlag, nil))
	}()

	select {}  
}
