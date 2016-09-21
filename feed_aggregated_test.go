package getstream_test

import (
	"encoding/json"
	"fmt"
	getstream "github.com/GetStream/stream-go"
	"github.com/pborman/uuid"
	"testing"
	"time"
)

func ExampleAggregatedFeed_AddActivity() {
	client, err := getstream.New(&getstream.Config{
		APIKey:    "APIKey",
		APISecret: "APISecret",
		AppID:     "AppID",
		Location:  "Region"})
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
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		t.Fatal(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Error(err)
	}

	err = PostTestCleanUp(client, nil, nil, []*getstream.Activity{activity})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAggregatedFeedAddActivityWithTo(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	toFeed, err := client.AggregatedFeed("aggregated", "barry")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
		To:        []getstream.Feed{toFeed},
	})
	if err != nil {
		t.Fatal(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = PostTestCleanUp(client, nil, nil, []*getstream.Activity{activity})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAggregatedFeedRemoveActivity(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:   "post",
		Object: getstream.FeedID("flat:eric"),
		Actor:  getstream.FeedID("flat:john"),
	})
	if err != nil {
		t.Fatal(err)
	}

	if activity.Verb != "post" {
		t.Fail()
	}

	rmActivity := getstream.Activity{
		ID: activity.ID,
	}

	err = feed.RemoveActivity(&rmActivity)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAggregatedFeedRemoveByForeignIDActivity(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "08f01c47-014f-11e4-aa8f-0cc47a024be0",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	PostTestCleanUp(client, nil, nil, []*getstream.Activity{activity})
}

func TestAggregatedFeedActivities(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	_, err = feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		t.Error(err)
	}

	activities, err := feed.Activities(&getstream.GetAggregatedFeedInput{})
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range activities.Results {
		err = PostTestCleanUp(client, nil, nil, result.Activities)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestAggregatedFeedAddActivities(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
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
		t.Error(err)
	}

	err = PostTestCleanUp(client, nil, nil, activities)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAggregatedFeedFollowUnfollow(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feedA, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	feedB, err := client.FlatFeed("flat", "eric")
	if err != nil {
		t.Fatal(err)
	}

	err = feedA.FollowFeedWithCopyLimit(feedB, 20)
	if err != nil {
		t.Fail()
	}

	// get feedB's followers, ensure feedA is there
	followers, err := feedB.FollowersWithLimitAndSkip(5, 0)
	if err != nil {
		t.Error(err)
	}
	if followers[0].UserID != "bob" {
		t.Error("Bob's aggregated feed is not a follower of FeedB")
	}

	// get things that feedA follows, ensure feedB is in there
	following, err := feedA.FollowingWithLimitAndSkip(5,0)
	if err != nil {
		t.Error(err)
	}
	if following[0].UserID != "eric" {
		t.Error("Eric's FeedB is not a follower of FeedA")
	}

	err = feedA.Unfollow(feedB)
	if err != nil {
		t.Fail()
	}

	PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB})
}

func TestAggregatedFeedFollowKeepingHistory(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feedA, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	feedB, err := client.FlatFeed("flat", "eric")
	if err != nil {
		t.Fatal(err)
	}

	err = feedA.FollowFeedWithCopyLimit(feedB, 20)
	if err != nil {
		t.Fail()
	}

	err = feedA.UnfollowKeepingHistory(feedB)
	if err != nil {
		t.Fail()
	}

	PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB})
}

func TestAggregatedFeedFollowingFollowers(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feedA, err := client.AggregatedFeed("aggregated", "bob")
	if err != nil {
		t.Fatal(err)
	}

	feedB, err := client.FlatFeed("flat", "eric")
	if err != nil {
		t.Fatal(err)
	}

	feedC, err := client.FlatFeed("flat", "barry")
	if err != nil {
		t.Fatal(err)
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

	PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB, feedC})
}

func TestAggregatedActivityMetaData(t *testing.T) {
	now := time.Now()

	data := struct {
		Foo  string
		Fooz string
	}{
		Foo:  "foo",
		Fooz: "fooz",
	}

	dataB, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	raw := json.RawMessage(dataB)

	activity := getstream.Activity{
		ForeignID: uuid.New(),
		Actor:     getstream.FeedID("user:eric"),
		Object:    getstream.FeedID("user:bob"),
		Target:    getstream.FeedID("user:john"),
		Origin:    getstream.FeedID("user:barry"),
		Verb:      "post",
		TimeStamp: &now,
		Data:      &raw,
		MetaData: map[string]string{
			"meta": "data",
		},
	}

	b, err := json.Marshal(&activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	b2, err := json.Marshal(activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	resultActivity := getstream.Activity{}
	err = json.Unmarshal(b, &resultActivity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resultActivity2 := getstream.Activity{}
	err = json.Unmarshal(b2, &resultActivity2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if resultActivity.ForeignID != activity.ForeignID {
		fmt.Println(activity.ForeignID)
		fmt.Println(resultActivity.ForeignID)
		t.Fail()
	}
	if resultActivity.Actor != activity.Actor {
		fmt.Println(activity.Actor)
		fmt.Println(resultActivity.Actor)
		t.Fail()
	}
	if resultActivity.Origin != activity.Origin {
		fmt.Println(activity.Origin)
		fmt.Println(resultActivity.Origin)
		t.Fail()
	}
	if resultActivity.Verb != activity.Verb {
		fmt.Println(activity.Verb)
		fmt.Println(resultActivity.Verb)
		t.Fail()
	}
	if resultActivity.Object != activity.Object {
		fmt.Println(activity.Object)
		fmt.Println(resultActivity.Object)
		t.Fail()
	}
	if resultActivity.Target != activity.Target {
		fmt.Println(activity.Target)
		fmt.Println(resultActivity.Target)
		t.Fail()
	}
	if resultActivity.TimeStamp.Format("2006-01-02T15:04:05.999999") != activity.TimeStamp.Format("2006-01-02T15:04:05.999999") {
		fmt.Println(activity.TimeStamp)
		fmt.Println(resultActivity.TimeStamp)
		t.Fail()
	}
	if resultActivity.MetaData["meta"] != activity.MetaData["meta"] {
		fmt.Println(activity.MetaData)
		fmt.Println(resultActivity.MetaData)
		t.Fail()
	}
	if string(*resultActivity.Data) != string(*activity.Data) {
		fmt.Println(string(*activity.Data))
		fmt.Println(string(*resultActivity.Data))
		t.Fail()
	}

	// fmt.Println(resultActivity)
	// fmt.Println(resultActivity.ForeignID)
	// fmt.Println(string(resultActivity.Data))
	// fmt.Println(resultActivity.MetaData)
	//
	// fmt.Println(resultActivity2)
	// fmt.Println(resultActivity2.ForeignID)
	// fmt.Println(string(resultActivity2.Data))
	// fmt.Println(resultActivity2.MetaData)

}
