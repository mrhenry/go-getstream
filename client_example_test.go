package getstream_test

import "fmt"
import "github.com/mrhenry/go-getstream"

func ExampleClient() {

	client := getstream.New("APIKey", "APISecret", "AppID", "Region")

	_ = client

}

func ExampleClient_FlatFeed() {

	client := getstream.New("APIKey", "APISecret", "AppID", "Region")

	feed, err := client.FlatFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = feed

}

func ExampleClient_NotificationFeed() {

	client := getstream.New("APIKey", "APISecret", "AppID", "Region")

	feed, err := client.NotificationFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = feed

}
