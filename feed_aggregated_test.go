package getstream

import (
	"fmt"
	"testing"
)

func ExampleAggregatedFeed_AddActivity() {

	opts := ServerOptions("APIKey", "APISecret", "AppID", "Region")
	client, err := New(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := client.AggregatedFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	activity, err := feed.AddActivity(&Activity{
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

func TestAggregatedFeedAddActivity(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	if feed == nil {
		fmt.Println("nil feed")
		t.Fail()
		return
	}

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
		return
	}

	err = testCleanUp(client, nil, nil, []*Activity{activity})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedAddActivityWithTo(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	toFeed, err := client.AggregatedFeed("aggregated", "barry")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
		To:        []Feed{toFeed},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
		return
	}

	err = testCleanUp(client, nil, nil, []*Activity{activity})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedRemoveActivity(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	activity, err := feed.AddActivity(&Activity{
		Verb:   "post",
		Object: FeedID("flat:eric"),
		Actor:  FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" {
		t.Fail()
		return
	}

	rmActivity := Activity{
		ID: activity.ID,
	}

	err = feed.RemoveActivity(&rmActivity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedRemoveByForeignIDActivity(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: "08f01c47-014f-11e4-aa8f-0cc47a024be0",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "08f01c47-014f-11e4-aa8f-0cc47a024be0" {
		t.Fail()
		return
	}

	rmActivity := Activity{
		ForeignID: activity.ForeignID,
	}
	_ = rmActivity

	err = feed.RemoveActivityByForeignID(activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	testCleanUp(client, nil, nil, []*Activity{activity})

}

func TestAggregatedFeedActivities(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	activities, err := feed.Activities(&GetAggregatedFeedInput{})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	for _, result := range activities.Results {
		err = testCleanUp(client, nil, nil, result.Activities)
		if err != nil {
			fmt.Println(err)
			t.Fail()
			return
		}
	}
}

func TestAggregatedFeedAddActivities(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	activities, err := feed.AddActivities([]*Activity{
		&Activity{
			Verb:      "post",
			ForeignID: "099978b6-3b72-4f5c-bc43-247ba6ae2dd9",
			Object:    FeedID("flat:eric"),
			Actor:     FeedID("flat:john"),
		}, &Activity{
			Verb:      "walk",
			ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
			Object:    FeedID("flat:john"),
			Actor:     FeedID("flat:eric"),
		},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	err = testCleanUp(client, nil, nil, activities)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedFollow(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.AggregatedFeed("aggregated", "bob")
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
		return
	}

	err = feedA.Unfollow(feedB)
	if err != nil {
		t.Fail()
		return
	}

	testCleanUpFollows(client, []*FlatFeed{feedB})

}

func TestAggregatedFeedFollowKeepingHistory(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.AggregatedFeed("aggregated", "bob")
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
		return
	}

	err = feedA.UnfollowKeepingHistory(feedB)
	if err != nil {
		t.Fail()
		return
	}

	testCleanUpFollows(client, []*FlatFeed{feedB})

}

func TestAggregatedFeedFollowingFollowers(t *testing.T) {

	client, err := testSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.AggregatedFeed("aggregated", "bob")
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

	feedC, err := client.FlatFeed("flat", "barry")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	err = feedA.FollowFeedWithCopyLimit(feedB, 20)
	if err != nil {
		t.Fail()
		return
	}

	err = feedA.FollowFeedWithCopyLimit(feedC, 20)
	if err != nil {
		t.Fail()
		return
	}

	_, err = feedA.FollowingWithLimitAndSkip(20, 0)
	if err != nil {
		t.Fail()
		return
	}

	testCleanUpFollows(client, []*FlatFeed{feedB, feedC})

}
