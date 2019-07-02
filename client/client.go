package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type HNClient struct {
	BaseURL string
	c       *colly.Collector
}

type FeedItem struct {
	Title    string `json:"title"`
	URI      string `json:"uri"`
	Author   string `json:"author"`
	Points   int    `json:"points"`
	Comments int    `json:"comments"`
	Rank     int    `json:"rank"`
}

//NewHNClient
func NewHNClient() HNClient {
	return HNClient{c: colly.NewCollector(), BaseURL: "https://news.ycombinator.com"}

}

func (hnc *HNClient) GetTopStories(targetAmount int) ([]*FeedItem, error) {
	if targetAmount == 0 {
		return nil, nil
	}
	pageNum := 1

	var allItems []*FeedItem

	for {
		items := hnc.scrapeMain(len(allItems), targetAmount, pageNum)
		allItems = append(allItems, items...)
		if len(allItems) != targetAmount {
			pageNum++
			//to ensure that there won't be an infinite loop
		} else if len(allItems) == 0 || len(items) == 0 || len(allItems) == targetAmount {
			break
		}

	}
	return allItems, nil

}

func (hnc *HNClient) scrapeMain(currentTotal int, maxStories int, pageNum int) []*FeedItem {
	var items []*FeedItem
	hnc.c.OnHTML("#hnmain .itemlist .athing", func(e *colly.HTMLElement) {
		if currentTotal == maxStories {
			return
		}
		item, err := hnc.processFeedItem(e)
		if err != nil {
			// log error, or write to file
			// fmt.Println(err)
		} else {
			items = append(items, item)
			currentTotal++

		}
	})
	hnc.c.Visit(fmt.Sprintf("%s/news?p=%d", hnc.BaseURL, pageNum))

	return items
}

//processFeedItem
func (hnc *HNClient) processFeedItem(e *colly.HTMLElement) (*FeedItem, error) {
	rank := e.ChildText(".rank")
	rankSplit := strings.Split(rank, ".")
	if len(rankSplit) == 0 {
		return nil, fmt.Errorf("rank expected to have . got %s", rank)
	}
	rankI, err := strconv.Atoi(rankSplit[0])
	if err != nil {
		return nil, fmt.Errorf("rank expected to be integer, got %s", rank)
	}
	title := e.ChildText(".title a")
	if len(title) > 256 {
		title = title[:256]
	}
	link := e.ChildAttr(".title a", "href")

	_, err = url.ParseRequestURI(link)
	if err != nil {
		return nil, fmt.Errorf("feed item does not contain a valid URI, got %s", link)
	}

	metaDataRow := e.DOM.Next()
	if metaDataRow == nil {
		return nil, fmt.Errorf("expected to have a metadata row, got none")
	}
	score := strings.TrimSpace(strings.Replace(metaDataRow.Find(".score").Text(), "points", "", -1))
	scoreI, err := strconv.Atoi(score)
	if err != nil {
		fmt.Sprintf("score expected to be integer, got %s", score)
	}
	author := metaDataRow.Find(".hnuser").Text()
	if len(author) > 256 {
		author = author[:256]
	}
	var comments string
	metaDataRow.Find("a").EachWithBreak((func(i int, s *goquery.Selection) bool {
		if strings.Contains(s.Text(), "comments") {
			comments = strings.TrimSpace(strings.Replace(s.Text(), "comments", "", -1))
			return false
		}
		return true
	}))
	commentsI, err := strconv.Atoi(comments)
	if err != nil {
		fmt.Sprintf("comments expected to be integer, got %s", comments)

	}
	return &FeedItem{title, link, author, scoreI, commentsI, rankI}, nil
}
