package getstream

import (
	"fmt"
	"testing"
)

func TestFlatFeedAddActivity(t *testing.T) {

	client, err := testSetup()
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

	err = testCleanUp(client, []*FlatFeedActivity{activity}, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedRemoveActivity(t *testing.T) {

	client, err := testSetup()
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

	rmActivity := FlatFeedActivity{
		ID: activity.ID,
	}

	err = feed.RemoveActivity(&rmActivity)
	if err != nil {
		t.Fail()
		return
	}
}

func TestFlatFeedRemoveByForeignIDActivity(t *testing.T) {

	client, err := testSetup()
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

	rmActivity := FlatFeedActivity{
		ForeignID: activity.ForeignID,
	}

	err = feed.RemoveActivityByForeignID(&rmActivity)
	if err != nil {
		t.Fail()
		return
	}
}

func TestFlatFeedListActivities(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		t.Fail()
		return
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fail()
		return
	}

	_, err = feed.AddActivity(&FlatFeedActivity{
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

	activities, err := feed.Activities(&GetFlatFeedInput{})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	err = testCleanUp(client, activities.Activities, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedAddActivities(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		t.Fail()
		return
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fail()
		return
	}

	activityA, err := feed.AddActivity(&FlatFeedActivity{
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

	activityB, err := feed.AddActivity(&FlatFeedActivity{
		Verb:      "walk",
		ForeignID: "088878b6-3b72-4f5c-bc43-247ba6ae2dd9",
		Object:    FeedID("flat:john"),
		Actor:     FeedID("flat:eric"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activityA.Verb != "post" && activityA.ForeignID != "099978b6-3b72-4f5c-bc43-247ba6ae2dd9" {
		t.Fail()
		return
	}

	if activityB.Verb != "walk" && activityB.ForeignID != "088878b6-3b72-4f5c-bc43-247ba6ae2dd9" {
		t.Fail()
		return
	}

	err = testCleanUp(client, []*FlatFeedActivity{activityA, activityB}, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}
