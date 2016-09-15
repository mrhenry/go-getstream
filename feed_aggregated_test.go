package getstream_test

import (
	"fmt"
	"testing"

	getstream "github.com/GetStream/stream-go"
)

func ExampleAggregatedFeed_AddActivity() {

	client, err := getstream.New("APIKey", "APISecret", "AppID", "Region")
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := client.AggregatedFeed("FeedSlug", "UserID")
	if err != nil {
		fmt.Println(err)
		return
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = activity
}

func TestAggregatedFeedAddActivity(t *testing.T) {

	client, err := getstream.PreTestSetup()
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

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = getstream.PostTestCleanUp(client, nil, nil, []*getstream.Activity{activity})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedAddActivityWithTo(t *testing.T) {

	client, err := getstream.PreTestSetup()
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

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
		To:        []getstream.Feed{toFeed},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = getstream.PostTestCleanUp(client, nil, nil, []*getstream.Activity{activity})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedRemoveActivity(t *testing.T) {

	client, err := getstream.PreTestSetup()
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

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:   "post",
		Object: getstream.FeedID("flat:eric"),
		Actor:  getstream.FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" {
		t.Fail()
	}

	rmActivity := getstream.Activity{
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

	client, err := getstream.PreTestSetup()
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

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "08f01c47-014f-11e4-aa8f-0cc47a024be0",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "08f01c47-014f-11e4-aa8f-0cc47a024be0" {
		t.Fail()
	}

	rmActivity := getstream.Activity{
		ForeignID: activity.ForeignID,
	}
	_ = rmActivity

	err = feed.RemoveActivityByForeignID(activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	getstream.PostTestCleanUp(client, nil, nil, []*getstream.Activity{activity})

}

func TestAggregatedFeedActivities(t *testing.T) {

	client, err := getstream.PreTestSetup()
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

	_, err = feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	activities, err := feed.Activities(&getstream.GetAggregatedFeedInput{})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	for _, result := range activities.Results {
		err  = getstream.PostTestCleanUp(client, nil, nil, result.Activities)
		if err != nil {
			fmt.Println(err)
			t.Fail()
			return
		}
	}
}

func TestAggregatedFeedAddActivities(t *testing.T) {

	client, err  := getstream.PreTestSetup()
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

	activities, err := feed.AddActivities([]*getstream.Activity{
		&getstream.Activity{
			Verb:      "post",
			ForeignID: "099978b6-3b72-4f5c-bc43-247ba6ae2dd9",
			Object:    getstream.FeedID("flat:eric"),
			Actor:     getstream.FeedID("flat:john"),
		}, &getstream.Activity{
			Verb:      "walk",
			ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
			Object:    getstream.FeedID("flat:john"),
			Actor:     getstream.FeedID("flat:eric"),
		},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err  = getstream.PostTestCleanUp(client, nil, nil, activities)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestAggregatedFeedFollow(t *testing.T) {

	client, err  := getstream.PreTestSetup()
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
	}

	err = feedA.Unfollow(feedB)
	if err != nil {
		t.Fail()
	}

	getstream.PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB})

}

func TestAggregatedFeedFollowKeepingHistory(t *testing.T) {

	client, err  := getstream.PreTestSetup()
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
	}

	err = feedA.UnfollowKeepingHistory(feedB)
	if err != nil {
		t.Fail()
	}

	getstream.PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB})

}

func TestAggregatedFeedFollowingFollowers(t *testing.T) {

	client, err  := getstream.PreTestSetup()
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
	}

	err = feedA.FollowFeedWithCopyLimit(feedC, 20)
	if err != nil {
		t.Fail()
	}

	_, err = feedA.FollowingWithLimitAndSkip(20, 0)
	if err != nil {
		t.Fail()
	}

	getstream.PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB, feedC})

}
