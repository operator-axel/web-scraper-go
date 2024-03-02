![webScraperLogo](https://github.com/operator-axel/web-scraper-go/assets/77085081/861bf66a-0747-478b-ac8a-d31fbb2f9a7a)

# Description
A web scraper built in Go, that asks for user input, domain, subdomain(optional), and the URL/link to scrape. 
The allowed domain input helps to keep the scraper contained to the specific content you want. 

The first prompt is for the 'Title' of the new note - as the scraper parses the links into a bulleted list of *Markdown Links* 
and creates a new note in my Obsidian vault. There is a comment above where the path/to/my/obsidian/vault is - change for your use case. 

# Quick Start 

`git clone https://github.com/operator-axel/web-scraper-go.git`

`cd web-scraper-go`

`go run *.go`
