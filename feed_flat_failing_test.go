package getstream_test

import (
	"fmt"
	"testing"
	"github.com/GetStream/stream-go"
)

func TestFlatFeedAddActivityFail(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "not a real foreign id",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err == nil {
		t.Fail()
		return
	}

	_, err = client.FlatFeed("flat&skinny", "bob@#awesome")
	if err == nil {
		t.Fail()
		return
	}

}
