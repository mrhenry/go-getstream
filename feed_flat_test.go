package getstream

import (
	"fmt"
	"testing"
)

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

func TestFlatFeedAddActivity(t *testing.T) {

	client, err := testSetup()
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

	activity, err := feed.AddActivity(&FlatFeedActivity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
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

	activity, err := feed.AddActivity(&FlatFeedActivity{
		Verb:   "post",
		Object: FeedID("flat:eric"),
		Actor:  FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if activity.Verb != "post" {
		t.Fail()
	}

	rmActivity := FlatFeedActivity{
		ID: activity.ID,
	}

	err = feed.RemoveActivity(&rmActivity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedRemoveByForeignIDActivity(t *testing.T) {

	client, err := testSetup()
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

	activity, err := feed.AddActivity(&FlatFeedActivity{
		Verb:      "post",
		ForeignID: "08f01c47-014f-11e4-aa8f-0cc47a024be0",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if activity.Verb != "post" && activity.ForeignID != "08f01c47-014f-11e4-aa8f-0cc47a024be0" {
		t.Fail()
	}

	rmActivity := FlatFeedActivity{
		ForeignID: activity.ForeignID,
	}
	_ = rmActivity

	err = feed.RemoveActivityByForeignID(activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	testCleanUp(client, []*FlatFeedActivity{activity}, nil)

}

func TestFlatFeedActivities(t *testing.T) {

	client, err := testSetup()
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

	_, err = feed.AddActivity(&FlatFeedActivity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	activities, err := feed.Activities(&GetFlatFeedInput{})
	if err != nil {
		fmt.Println(err)
		t.Fail()
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

	activities, err := feed.AddActivities([]*FlatFeedActivity{
		&FlatFeedActivity{
			Verb:      "post",
			ForeignID: "099978b6-3b72-4f5c-bc43-247ba6ae2dd9",
			Object:    FeedID("flat:eric"),
			Actor:     FeedID("flat:john"),
		}, &FlatFeedActivity{
			Verb:      "walk",
			ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
			Object:    FeedID("flat:john"),
			Actor:     FeedID("flat:eric"),
		},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = testCleanUp(client, activities, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedFollow(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.FlatFeed("flat", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedB, err := client.FlatFeed("flat", "eric")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	err = feedA.FollowFeedWithCopyLimit(feedB, 20)
	if err != nil {
		t.Fail()
	}

	testCleanUpFollows(client, []*FlatFeed{feedA, feedB})

}
