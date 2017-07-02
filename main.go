package main

import (
	"flag"
	"github.com/sdorunga1/gocrawl/client"
	"github.com/sdorunga1/gocrawl/console"
	"github.com/sdorunga1/gocrawl/crawler"
)

var (
	startPage = flag.String("s", "", "Url to use as the location to start crawling from, in the format `http://example.com")
)

func main() {
	flag.Parse()

	engine := crawler.NewCrawler(client.NewClient(), &console.TerminalPrinter{})
	engine.Crawl(*startPage)
}
