package html

import (
	"reflect"
	"testing"
)

const (
	twoUrls = `
  <!DOCTYPE html>
  <html>
    <body>
      <a href="https://site.com/about">About</a>
      <a href="/blog">Blog</a>
      <a href="https://externalsite.com/about">About External</a>
    </body>
  </html> `
	assetUrls = `
  <!DOCTYPE html>
  <html>
    <head>
      <script src="myscript.js"></script>
      <link rel="stylesheet" type="text/css" href="mystyle.css">
    </head>
    <body>
      <img src="https://site.com/img/horse.png"/>
    </body>
  </html>`
	duplicates = `
  <!DOCTYPE html>
  <html>
    <head>
      <script src="myscript.js"></script>
      <script src="myscript.js"></script>
    </head>
    <body>
      <img src="https://site.com/img/horse.png"/>
      <img src="https://site.com/img/horse.png"/>
      <a href="https://site.com/about">About</a>
      <a href="https://site.com/about">About</a>
    </body>
  </html>`
)

func TestReturnsOnlyPagesForCrawledSite(t *testing.T) {
	pageUrl := "https://site.com"
	html := twoUrls
	page := Page{html, pageUrl}
	links := page.Links()

	expectedLinks := []string{"https://site.com/about", "https://site.com/blog"}

	if !reflect.DeepEqual(links.PageUrls, expectedLinks) {
		t.Error("Expected: ", expectedLinks, ". Got: ", links)
	}
}

func TestReturnsAListOfAssetLinksFromHttp(t *testing.T) {
	pageUrl := "https://site.com"
	html := assetUrls
	page := Page{html, pageUrl}
	links := page.Links()
	expectedLinks := []string{"https://site.com/myscript.js", "https://site.com/mystyle.css", "https://site.com/img/horse.png"}
	if !reflect.DeepEqual(links.AssetUrls, expectedLinks) {
		t.Error("Expected: ", expectedLinks, ". Got: ", links)
	}
}

func TestFiltersOutDuplicateUrls(t *testing.T) {
	pageUrl := "https://site.com"
	html := duplicates
	page := Page{html, pageUrl}
	links := page.Links()
	expectedPageLinks := []string{"https://site.com/about"}
	expectedAssetLinks := []string{"https://site.com/myscript.js", "https://site.com/img/horse.png"}
	if !reflect.DeepEqual(links.PageUrls, expectedPageLinks) {
		t.Error("Expected: ", expectedPageLinks, ". Got: ", links.PageUrls)
	}

	if !reflect.DeepEqual(links.AssetUrls, expectedAssetLinks) {
		t.Error("Expected: ", expectedAssetLinks, ". Got: ", links.AssetUrls)
	}
}
