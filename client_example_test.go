package getstream_test

import (
	"fmt"

	getstream "github.com/mrhenry/go-getstream"
)

func ExampleClient() {

	opts := getstream.ServerOptions("APIKey", "APISecret", "AppID", "Region")
	client, err := getstream.New(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = client

}

func ExampleClient_FlatFeed() {

	opts := getstream.ServerOptions("APIKey", "APISecret", "AppID", "Region")
	client, err := getstream.New(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := client.FlatFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = feed

}

func ExampleClient_NotificationFeed() {

	opts := getstream.ServerOptions("APIKey", "APISecret", "AppID", "Region")
	client, err := getstream.New(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := client.NotificationFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = feed

}
