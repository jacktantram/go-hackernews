package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/go-hackernews/client"
)

func main() {
	var posts = flag.Int("posts", 1, "how many posts to print from hackernews")
	flag.Parse()
	client := client.NewHNClient()
	pp, err := client.GetTopStories(*posts)
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := json.Marshal(pp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))

}
