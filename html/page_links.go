package html

import (
  "fmt"
  "golang.org/x/net/html"
  "net/url"
  "strings"
)

type Page struct {
	Html string
	Url  string
}

type PageLinks struct {
	AssetUrls []string
	PageUrls  []string
	OriginUrl string
}

var pageTags = []string{"a"}
var assetTags = []string{"link", "script", "img"}

func (page *Page) Links() PageLinks {
	parsedBody, err := html.Parse(strings.NewReader(page.Html))
	if err != nil {
		return PageLinks{}
	}
	pageLinks := page.collectLinks(parsedBody, []string{}, pageTags)
	resourceLinks := page.collectLinks(parsedBody, []string{}, assetTags)

	return PageLinks{resourceLinks, pageLinks, page.Url}
}

func (page *Page) collectLinks(rootNode *html.Node, links []string, types []string) []string {
	for node := rootNode.FirstChild; node != nil; node = node.NextSibling {
		if link := page.extractInternalLink(node, types); link != "" {
			if !contains(links, link) {
				links = append(links, link)
			}
		}
		links = page.collectLinks(node, links, types)
	}

	return links
}

func (page *Page) extractInternalLink(node *html.Node, types []string) (link string) {
	if node.Type == html.ElementNode && contains(types, node.Data) {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" || attribute.Key == "src" {
				if absoluteLink := page.absolutifyUrl(attribute.Val); page.isInternalLink(absoluteLink) {
					link = page.absolutifyUrl(attribute.Val).String()
				}
			}
		}
	}
	return
}

func contains(list []string, item string) bool {
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}
	return false
}

func (page *Page) absolutifyUrl(link string) *url.URL {
	linkUrl, err := url.Parse(link)
	if err != nil {
		fmt.Println("Malformed Url", link)
	}
	if linkUrl.Host == "" {
		linkUrl.Scheme = page.parsedUrl().Scheme
		linkUrl.Host = page.parsedUrl().Host
	}
	return linkUrl
}

func (page *Page) parsedUrl() *url.URL {
	parsedUrl, err := url.Parse(page.Url)
	if err != nil {
		fmt.Println("Malformed Url", page.Url)
	}
	return parsedUrl
}

func (page *Page) isInternalLink(link *url.URL) bool {
	return link.Host == page.parsedUrl().Host
}
