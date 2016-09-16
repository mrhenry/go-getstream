package getstream_test

import (
	"github.com/GetStream/stream-go"
	"testing"
)

func TestFlatFeedAddActivityFail(t *testing.T) {

	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fatal(err)
	}

	_, err = feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "not a real foreign id",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err == nil {
		t.Fatal(err)
	}

	_, err = client.FlatFeed("flat&skinny", "bob@#awesome")
	if err == nil {
		t.Fatal(err)
	}
}
