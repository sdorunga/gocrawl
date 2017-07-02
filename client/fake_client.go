package client

import (
)

type FakeClient struct {
  sites map[string]string
}

func (this *FakeClient) Get(url string) string {
  site, present := this.sites[url]
  if !present {
    panic("Couldn't find " + url + " in list. Make sure you added the url")
  }
  return site
}

func (this *FakeClient) StubAndReturn(url string, html string) {
  this.sites[url] = html
}

func NewFakeClient() *FakeClient {
	return &FakeClient{make(map[string]string)}
}
