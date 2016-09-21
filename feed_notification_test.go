package getstream_test

import (
	"encoding/json"
	"fmt"
	"github.com/GetStream/stream-go"
	"testing"
	"github.com/pborman/uuid"
	"time"
)

func ExampleNotificationFeed_AddActivity() {
	client, err := getstream.New(&getstream.Config{
		APIKey:    "APIKey",
		APISecret: "APISecret",
		AppID:     "AppID",
		Location:  "Region"})
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := client.NotificationFeed("FeedSlug", "UserID")
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

func TestNotificationFeedAddActivity(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "bob")
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
		t.Error(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = PostTestCleanUp(client, nil, []*getstream.Activity{activity}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNotificationFeedAddActivityWithTo(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "bob")
	if err != nil {
		t.Fatal(err)
	}

	feedTo, err := client.NotificationFeed("notification", "barry")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
		To:        []getstream.Feed{feedTo},
	})
	if err != nil {
		t.Error(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = PostTestCleanUp(client, nil, []*getstream.Activity{activity}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNotificationFeedRemoveActivity(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "bob")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:   "post",
		Object: getstream.FeedID("flat:eric"),
		Actor:  getstream.FeedID("flat:john"),
	})
	if err != nil {
		t.Error(err)
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

func TestNotificationFeedRemoveByForeignIDActivity(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "bob")
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
		t.Error(err)
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

	PostTestCleanUp(client, nil, []*getstream.Activity{activity}, nil)
}

func TestNotificationFeedActivities(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "bob")
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

	activities, err := feed.Activities(&getstream.GetNotificationFeedInput{})
	if err != nil {
		t.Error(err)
	}

	for _, result := range activities.Results {
		err = PostTestCleanUp(client, nil, result.Activities, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNotificationFeedAddActivities(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "bob")
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

	err = PostTestCleanUp(client, nil, activities, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNotificationFeedFollow(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feedA, err := client.NotificationFeed("notification", "bob")
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

	err = feedA.Unfollow(feedB)
	if err != nil {
		t.Fail()
	}

	PostTestCleanUpFollows(client, []*getstream.FlatFeed{feedB})
}

func TestNotificationFeedFollowKeepingHistory(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feedA, err := client.NotificationFeed("notification", "bob")
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

func TestNotificationFeedFollowingFollowers(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feedA, err := client.NotificationFeed("notification", "bob")
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

func TestMarkAsSeen(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "larry")
	if err != nil {
		t.Fatal(err)
	}

	feed.AddActivities([]*getstream.Activity{
		&getstream.Activity{
			Actor:  getstream.FeedID("flat:larry"),
			Object: getstream.FeedID("notification:larry"),
			Verb:   "post",
		},
	})

	time.Sleep(time.Second * 2)

	output, _ := feed.Activities(nil)
	if output.Unseen == 0 {
		t.Fail()
	}

	feed.MarkActivitiesAsSeenWithLimit(15)

	time.Sleep(time.Second * 2)

	output, _ = feed.Activities(nil)
	if output.Unseen != 0 {
		t.Fail()
	}

	for _, result := range output.Results {
		err = PostTestCleanUp(client, nil, result.Activities, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestMarkAsRead(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("notification", "larry")
	if err != nil {
		t.Fatal(err)
	}

	feed.AddActivities([]*getstream.Activity{
		&getstream.Activity{
			Actor:  getstream.FeedID("flat:larry"),
			Object: getstream.FeedID("notification:larry"),
			Verb:   "post",
		},
	})

	time.Sleep(time.Second * 2)

	output, _ := feed.Activities(nil)
	if output.Unread == 0 {
		t.Fail()
	}

	for _, result := range output.Results {
		err = feed.MarkActivitiesAsRead(result.Activities)
		if err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(time.Second * 2)

	output, _ = feed.Activities(nil)
	if output.Unread != 0 {
		t.Fail()
	}

	for _, result := range output.Results {
		err = PostTestCleanUp(client, nil, result.Activities, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNotificationActivityMetaData(t *testing.T) {

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
