package getstream_test

import (
	"fmt"
	"testing"
	"time"
	"github.com/GetStream/stream-go"
)

func ExampleNotificationFeed_AddActivity() {

	client, err := getstream.New("APIKey", "APISecret", "AppID", "Region")
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

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "bob")
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
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = getstream.PostTestCleanUp(client, nil, []*getstream.Activity{activity}, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestNotificationFeedAddActivityWithTo(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedTo, err := client.NotificationFeed("notification", "barry")
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
		To:        []getstream.Feed{feedTo},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = getstream.PostTestCleanUp(client, nil, []*getstream.Activity{activity}, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestNotificationFeedRemoveActivity(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "bob")
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

func TestNotificationFeedRemoveByForeignIDActivity(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "bob")
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

	getstream.PostTestCleanUp(client, nil, []*getstream.Activity{activity}, nil)

}

func TestNotificationFeedActivities(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "bob")
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

	activities, err := feed.Activities(&getstream.GetNotificationFeedInput{})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, result := range activities.Results {
		err = getstream.PostTestCleanUp(client, nil, result.Activities, nil)
		if err != nil {
			fmt.Println(err)
			t.Fail()
			return
		}
	}
}

func TestNotificationFeedAddActivities(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "bob")
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

	err = getstream.PostTestCleanUp(client, nil, activities, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestNotificationFeedFollow(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.NotificationFeed("notification", "bob")
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

func TestNotificationFeedFollowKeepingHistory(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.NotificationFeed("notification", "bob")
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

func TestNotificationFeedFollowingFollowers(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedA, err := client.NotificationFeed("notification", "bob")
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

func TestMarkAsSeen(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "larry")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
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
		err = getstream.PostTestCleanUp(client, nil, result.Activities, nil)
		if err != nil {
			fmt.Println(err)
			t.Fail()
			return
		}
	}
}

func TestMarkAsRead(t *testing.T) {

	client, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := client.NotificationFeed("notification", "larry")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
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
			fmt.Println(err)
			t.Fail()
			return
		}
	}

	time.Sleep(time.Second * 2)

	output, _ = feed.Activities(nil)
	if output.Unread != 0 {
		t.Fail()
	}

	for _, result := range output.Results {
		err = getstream.PostTestCleanUp(client, nil, result.Activities, nil)
		if err != nil {
			fmt.Println(err)
			t.Fail()
			return
		}
	}
}
