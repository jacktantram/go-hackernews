package main

import (
	"fmt"

	"github.com/go-hackernews/client"
)

func main() {
	client := client.NewHNClient()
	pp, _ := client.GetTopStories(101)

	fmt.Println(fmt.Sprintf("%v", pp))

}
