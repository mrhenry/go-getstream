package getstream

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pborman/uuid"
)

func TestExampleNotificationFeed_AddActivity(t *testing.T) {
	client, err := PreTestSetup()
	if err != nil {
		t.Fatal(err)
	}

	feed, err := client.NotificationFeed("flat", "UserID")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: uuid.New(),
		Object:    "flat:eric",
		Actor:     "flat:john",
	})
	if err != nil {
		t.Fatal(err)
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

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: uuid.New(),
		Object:    "flat:eric",
		Actor:     "flat:john",
	})
	if err != nil {
		t.Error(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = PostTestCleanUp(client, nil, []*Activity{activity}, nil)
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

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: uuid.New(),
		Object:    "flat:eric",
		Actor:     "flat:john",
		To:        []Feed{feedTo},
	})
	if err != nil {
		t.Error(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = PostTestCleanUp(client, nil, []*Activity{activity}, nil)
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

	activity, err := feed.AddActivity(&Activity{
		Verb:   "post",
		Object: "flat:eric",
		Actor:  "flat:john",
	})
	if err != nil {
		t.Error(err)
	}

	if activity.Verb != "post" {
		t.Fail()
	}

	rmActivity := Activity{
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

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: uuid.New(),
		Object:    "flat:eric",
		Actor:     "flat:john",
	})
	if err != nil {
		t.Error(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "08f01c47-014f-11e4-aa8f-0cc47a024be0" {
		t.Fail()
	}

	rmActivity := Activity{
		ForeignID: activity.ForeignID,
	}
	_ = rmActivity

	err = feed.RemoveActivityByForeignID(activity)
	if err != nil {
		t.Fatal(err)
	}

	PostTestCleanUp(client, nil, []*Activity{activity}, nil)
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

	_, err = feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: uuid.New(),
		Object:    "flat:eric",
		Actor:     "flat:john",
	})
	if err != nil {
		t.Error(err)
	}

	activities, err := feed.Activities(&GetNotificationFeedInput{})
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

	activities, err := feed.AddActivities([]*Activity{
		&Activity{
			Verb:      "post",
			ForeignID: uuid.New(),
			Object:    "flat:eric",
			Actor:     "flat:john",
		}, &Activity{
			Verb:      "walk",
			ForeignID: uuid.New(),
			Object:    "flat:john",
			Actor:     "flat:eric",
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

	PostTestCleanUpFlatFeedFollows(client, []*FlatFeed{feedB})
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

	PostTestCleanUpFlatFeedFollows(client, []*FlatFeed{feedB})
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

	PostTestCleanUpFlatFeedFollows(client, []*FlatFeed{feedB, feedC})
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

	feed.AddActivities([]*Activity{
		&Activity{
			Actor:  "flat:larry",
			Object: "notification:larry",
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

	feed.AddActivities([]*Activity{
		&Activity{
			Actor:  "flat:larry",
			Object: "notification:larry",
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
		t.Fatal(err)
	}

	raw := json.RawMessage(dataB)

	activity := Activity{
		ForeignID: uuid.New(),
		Actor:     "user:eric",
		Object:    "user:bob",
		Target:    "user:john",
		Origin:    "user:barry",
		Verb:      "post",
		TimeStamp: &now,
		Data:      &raw,
		MetaData: map[string]interface{}{
			"meta": "data",
		},
	}

	b, err := json.Marshal(&activity)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := json.Marshal(activity)
	if err != nil {
		t.Fatal(err)
	}

	resultActivity := Activity{}
	err = json.Unmarshal(b, &resultActivity)
	if err != nil {
		t.Fatal(err)
	}

	resultActivity2 := Activity{}
	err = json.Unmarshal(b2, &resultActivity2)
	if err != nil {
		t.Fatal(err)
	}

	if resultActivity.ForeignID != activity.ForeignID {
		t.Error(activity.ForeignID, resultActivity.ForeignID)
	}
	if resultActivity.Actor != activity.Actor {
		t.Error(activity.Actor, resultActivity.Actor)
	}
	if resultActivity.Origin != activity.Origin {
		t.Error(activity.Origin, resultActivity.Origin)
	}
	if resultActivity.Verb != activity.Verb {
		t.Error(activity.Verb, resultActivity.Verb)
	}
	if resultActivity.Object != activity.Object {
		t.Error(activity.Object, resultActivity.Object)
	}
	if resultActivity.Target != activity.Target {
		t.Error(activity.Target, resultActivity.Target)
	}
	if resultActivity.TimeStamp.Format("2006-01-02T15:04:05.999999") != activity.TimeStamp.Format("2006-01-02T15:04:05.999999") {
		t.Error(activity.TimeStamp, resultActivity.TimeStamp)
	}
	if resultActivity.MetaData["meta"] != activity.MetaData["meta"] {
		t.Error(activity.MetaData, resultActivity.MetaData)
	}
	if string(*resultActivity.Data) != string(*activity.Data) {
		t.Error(string(*activity.Data), string(*resultActivity.Data))
	}
}
