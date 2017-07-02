package crawler

import (
	"testing"
        "github.com/sdorunga1/gocrawl/client"
        "github.com/sdorunga1/gocrawl/console"
)

const (
  homepage = `
  <!DOCTYPE html>
  <html>
    <body>
      <a href="http://site.com/homepage">Home</a>
      <a href="/linked_page">Home</a>
      <a href="http://externalsite.com/about">About External</a>
      <img src="/site_logo.png"/>
    </body>
  </html> `
  linkedPage = `
  <!DOCTYPE html>
  <html>
    <body>
      <a href="http://site.com/homepage">Home</a>
      <img src="/site_logo.png"/>
    </body>
  </html> `
  expectedCrawlOutput = `
http://site.com/homepage
 Linked Pages:
   http://site.com/homepage
   http://site.com/linked_page
 Linked Assets:
   http://site.com/site_logo.png
http://site.com/linked_page
 Linked Pages:
   http://site.com/homepage
 Linked Assets:
   http://site.com/site_logo.png`
)

func TestCrawlerPrintsAssetAndUrls(t *testing.T) {
  fakePrinter := console.FakePrinter{}
  fakeClient := client.NewFakeClient()
  fakeClient.StubAndReturn("http://site.com/homepage", homepage)
  fakeClient.StubAndReturn("http://site.com/linked_page", linkedPage)
  crawler := NewCrawler(fakeClient, &fakePrinter)
  crawler.Crawl("http://site.com/homepage")
  if fakePrinter.Verify(expectedCrawlOutput) != true {
    t.Error("Expected: \n", expectedCrawlOutput, "\nGot: \n", fakePrinter.Messages)
  }
}
