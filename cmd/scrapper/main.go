package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"goWebScrapper/internal/scrapper"

	"github.com/sirupsen/logrus"
)

var (
	urlFlag       = flag.String("url", "", "The URL to scrape content from")
	outputFlag    = flag.String("output", "output", "The filename (without extension) to save the content")
	formatFlag    = flag.String("format", "txt", "The format to save the content")
	portFlag      = flag.String("port", "", "The port to serve the file for download")
	selectorsFlag = flag.String("selectors", "", "Comma-separated list of selector elements to scrape")
	concurrency   = flag.Int("concurrency", 5, "Number of concurrent scrapers")
	rateLimit     = flag.Int("rate-limit", 5, "Number of requests per second")
)

func main() {
	flag.Parse()

	if *urlFlag == "" {
		log.Fatal("Error: Please provide a URL using the -url flag")
		flag.Usage()
		os.Exit(1)
	}

	if *portFlag == "" {
		log.Fatal("Error: Please provide a TCP port to download the scraped content")
		flag.Usage()
		os.Exit(1)
	}

	if *selectorsFlag == "" {
		log.Fatal("Error: Please provide selectors using the -selectors flag")
		flag.Usage()
		os.Exit(1)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	ticker := time.NewTicker(time.Second / time.Duration(*rateLimit))
	defer ticker.Stop()

	urls := strings.Split(*urlFlag, ",")

	var wg sync.WaitGroup
	workerPool := make(chan struct{}, *concurrency)   

	for _, url := range urls {
		wg.Add(1)
		workerPool <- struct{}{}  
		go func(url string) {
			defer wg.Done()
			defer func() { <-workerPool }()  
			<-ticker.C

			doc, err := scrapper.Fetch(url)
			if err != nil {
				logrus.Errorf("Error fetching URL %s: %v", url, err)
				return
			}

			data, err := scrapper.ParseHTML(doc, *selectorsFlag)
			if err != nil {
				logrus.Errorf("Error parsing HTML for URL %s: %v", url, err)
				return
			}

			filename := scrapper.DetermineFileFormat(*formatFlag, *outputFlag, data)
			logrus.Infof("Content successfully scraped from %s and saved to %s", url, filename)
		}(url)
	}

	wg.Wait()
	close(workerPool)  

	http.Handle("/", http.FileServer(http.Dir(".")))
	logrus.Infof("Download your file at: http://localhost:%s/%s", *portFlag, *outputFlag)
	log.Fatal(http.ListenAndServe(":"+*portFlag, nil))
}
