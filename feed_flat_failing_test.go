package getstream

import (
	"fmt"
	"testing"
)

func TestFlatFeedAddActivityFail(t *testing.T) {

	client, err := testSetupRealClient()
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

	_, err = feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: "not a real foreign id",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
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
