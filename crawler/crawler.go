package crawler

import (
	"github.com/sdorunga1/gocrawl/client"
	"github.com/sdorunga1/gocrawl/console"
	"github.com/sdorunga1/gocrawl/collection"
	"github.com/sdorunga1/gocrawl/html"
	"sync"
)

type Crawler struct {
	Client client.HttpClient
	Printer console.Printer
}

func NewCrawler(client client.HttpClient, printer console.Printer) Crawler {
	return Crawler{client, printer}
}

func (crawler *Crawler) Crawl(startingUrl string) {
	wg := &sync.WaitGroup{}
	linksToCrawl := make(chan string, 10)
	results := make(chan html.PageLinks)
	seen := collection.NewSet()

	wg.Add(1)
	go func() { linksToCrawl <- startingUrl }()
	go crawler.crawlLinks(seen, linksToCrawl, wg, results)
	go crawler.printResults(results, wg)
	wg.Wait()
}

func (crawler *Crawler) crawlLinks(seen collection.StringSet, linksToCrawl chan string, wg *sync.WaitGroup, results chan html.PageLinks) {
	for link := range linksToCrawl {
		if seen.Contains(link) {
			continue
		}
		go func(link string) {
			seen.Add(link)
			basePage := html.Page{crawler.Client.Get(link), link}
			links := basePage.Links()
			results <- links

			for _, subLink := range links.PageUrls {
				if !seen.Contains(subLink) {
					wg.Add(1)
					linksToCrawl <- subLink
				}
			}
		}(link)
	}
}

func (crawler *Crawler) printResults(results chan html.PageLinks, wg *sync.WaitGroup) {
	for result := range results {
                crawler.print(result.OriginUrl)
                crawler.print(" " + "Linked Pages:")
		for _, pageUrl := range result.PageUrls {
                        crawler.print("   " + pageUrl)
		}
                crawler.print(" " + "Linked Assets:")
		for _, assetUrl := range result.AssetUrls {
                        crawler.print("   " + assetUrl)
		}
		wg.Done()
	}
}

func (crawler *Crawler) print(message string) {
  crawler.Printer.Print(message)
}
