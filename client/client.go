package client

import (
	"github.com/gocolly/colly"
)

type HNClient struct {
	BaseURL string
}

func NewHNClient() HNClient {
	c := colly.NewCollector()

}

type Item struct {
	Title    string
	URI      string
	Author   string
	Points   string
	Comments int
	Rank     string
}

func (hnc *HNClient) GetTopStories() []Item {
	return []Item{}
}

func (hnc *HNClient) GetItem() Item {
	return Item{}
}
