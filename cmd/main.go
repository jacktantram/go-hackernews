package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-hackernews/client"
)

func main() {
	client := client.NewHNClient()
	pp, _ := client.GetTopStories(100)

	_, _ = json.Marshal(pp)
	fmt.Println(len(pp))

}
