package getstream

import "fmt"

func ExampleFlatFeed_AddActivity() {

	client, err := New("APIKey", "APISecret", "AppID", "Region")
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := client.FlatFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	activity, err := feed.AddActivity(&FlatFeedActivity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = activity
}
