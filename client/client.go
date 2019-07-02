package client

type HNClient struct {
	BaseURL string
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
