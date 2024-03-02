package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var fileName string
	fmt.Print("What should the name of the file be? ")
	fileName, _ = reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	filePath := fmt.Sprintf("/home/ch40s/Obsidian/dolos/%s.md", fileName)
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", filePath, err)
		return
	}
	defer file.Close()

	var domain string
	fmt.Print("Provide the allowed domains you want to scrape: (e.g., archwiki.com) ")
	domain, _ = reader.ReadString('\n')
	domain = strings.TrimSpace(domain)

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	fmt.Print("Enter a substring that must be present in the links (leave blank for no filter): ")
	linkFilter, _ := reader.ReadString('\n')
	linkFilter = strings.TrimSpace(linkFilter)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Process links that contain the specified filter substring
		if linkFilter == "" || strings.Contains(link, linkFilter) {
			absoluteURL := e.Request.AbsoluteURL(link)

			parts := strings.Split(link, "/")
			var linkText string
			for i := len(parts) - 1; i >= 0; i-- {
				if parts[i] != "" {
					linkText = parts[i]
					break
				}
			}
			mdLink := "- [" + linkText + "](" + absoluteURL + ")\n"

			if _, err := file.WriteString(mdLink); err != nil {
				log.Printf("Failed to write lesson URL to file: %s\n", err)
			}
		}
	})

	var urlToScrape string
	fmt.Print("URL to scrape: ")
	fmt.Scan(&urlToScrape)
	err = c.Visit(urlToScrape)

	if err != nil {
		log.Fatalf("Failed to visit site: %s\n", err)
		return
	}

	log.Printf("Scraping finished, check file %q for results\n", filePath)
}
