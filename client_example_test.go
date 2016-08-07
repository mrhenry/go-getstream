package getstream_test

import "fmt"
import "github.com/mrhenry/go-getstream"

func ExampleClient() {

	client, err := getstream.New("APIKey", "APISecret", "AppID", "Region")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = client

}

func ExampleClient_FlatFeed() {

	client, err := getstream.New("APIKey", "APISecret", "AppID", "Region")
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

	client, err := getstream.New("APIKey", "APISecret", "AppID", "Region")
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
