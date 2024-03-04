package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Print("URL to scrape: ")
	var urlToScrape string
	fmt.Scan(&urlToScrape)

	parsedURL, err := url.Parse(urlToScrape)
	if err != nil {
		log.Fatalf("Failed to parse URL: %s\n", err)
	}

	// Automatically determine the link filter from the URL path
	// Extracting the first path segment as the link filter
	pathSegments := strings.Split(parsedURL.Path, "/")
	var linkFilter string
	if len(pathSegments) > 1 {
		linkFilter = "/" + pathSegments[1] // Adjust this logic based on your needs
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What should the name of the file be? ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	filePath := fmt.Sprintf("/home/ch40s/Obsidian/dolos/%s.md", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", filePath, err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(
		colly.AllowedDomains(parsedURL.Hostname()),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if strings.Contains(link, linkFilter) {
			absoluteURL := e.Request.AbsoluteURL(link)
			mdLink := fmt.Sprintf("- [%s](%s)\n", link, absoluteURL)
			if _, err := file.WriteString(mdLink); err != nil {
				log.Printf("Failed to write link to file: %s\n", err)
			}
		}
	})

	err = c.Visit(urlToScrape)
	if err != nil {
		log.Fatalf("Failed to visit site: %s\n", err)
		return
	}

	log.Printf("Scraping finished, check file %q for results\n", filePath)
}
