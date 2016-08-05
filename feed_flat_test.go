package getstream

import (
	"fmt"
	"os"
	"testing"
)

func TestFlatFeedAddActivity(t *testing.T) {

	client, err := testFlatFeedSetup()
	if err != nil {
		t.Fail()
		return
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fail()
		return
	}

	activity, err := feed.AddActivity(&FlatFeedActivity{
		Verb:      "post",
		ForeignID: "099978b6-3b72-4f5c-bc43-247ba6ae2dd9",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "099978b6-3b72-4f5c-bc43-247ba6ae2dd9" {
		t.Fail()
		return
	}

	err = testFlatFeedCleanUp(client, []*FlatFeedActivity{activity}, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func testFlatFeedSetup() (*Client, error) {

	testAPIKey := os.Getenv("key")
	testAPISecret := os.Getenv("secret")
	testAppID := os.Getenv("app_id")
	testRegion := os.Getenv("region")

	client, err := New(testAPIKey, testAPISecret, testAppID, testRegion)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func testFlatFeedCleanUp(client *Client, flats []*FlatFeedActivity, notifications []*NotificationFeedActivity) error {

	if len(flats) > 0 {

		feed, err := client.FlatFeed("flat", "bob")
		if err != nil {
			return err
		}

		for _, activity := range flats {
			err := feed.RemoveActivity(activity)
			if err != nil {
				return err
			}
		}
	}

	if len(notifications) > 0 {

		feed, err := client.NotificationFeed("notification", "bob")
		if err != nil {
			return err
		}

		for _, activity := range notifications {
			err := feed.RemoveActivity(activity)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
