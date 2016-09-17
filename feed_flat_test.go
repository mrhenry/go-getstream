package getstream

import (
	"fmt"
	"testing"

	"github.com/pborman/uuid"
)

func TestFlatFeedAddActivity(t *testing.T) {

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

	activity, err := feed.AddActivity(&Activity{
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

	err = testCleanUp(client, []*Activity{activity}, nil, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedAddActivityWithTo(t *testing.T) {

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

	feedTo, err := client.FlatFeed("flat", "barry")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feedToB, err := client.FlatFeed("flat", "larry")
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
		To:        []Feed{feedTo, feedToB},
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	err = testCleanUp(client, []*Activity{activity}, nil, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedUUID(t *testing.T) {

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

	var activities []*Activity

	for i := 0; i < 10; i++ {

		foreignID := uuid.New()

		activity, err := feed.AddActivity(&Activity{
			Verb:      "post",
			ForeignID: foreignID,
			Object:    FeedID("flat:eric"),
			Actor:     FeedID("flat:john"),
		})
		if err != nil {
			t.Log("fail add activity with UUID : ")
			t.Log(err)
			continue
		}

		err = feed.RemoveActivityByForeignID(activity)
		if err != nil {
			t.Log("fail remove activity with UUID : ")
			t.Log(err)
		}

		activities = append(activities, activity)
	}

	testCleanUp(client, activities, nil, nil)
}

func TestFlatFeedRemoveActivity(t *testing.T) {

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

	activity, err := feed.AddActivity(&Activity{
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

func TestFlatFeedRemoveByForeignIDActivity(t *testing.T) {

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

	activity, err := feed.AddActivity(&Activity{
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

	rmActivity := Activity{
		ForeignID: activity.ForeignID,
	}
	_ = rmActivity

	err = feed.RemoveActivityByForeignID(activity)
	if err != nil {
		fmt.Println(err)
		return
	}

	testCleanUp(client, []*Activity{activity}, nil, nil)

}

func TestFlatFeedActivities(t *testing.T) {

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

	err = testCleanUp(client, activities.Activities, nil, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedAddActivities(t *testing.T) {

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
	}

	err = testCleanUp(client, activities, nil, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestFlatFeedFollow(t *testing.T) {

	client, err := testSetupRealClient()
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

func TestFlatFeedFollowingFollowers(t *testing.T) {

	client, err := testSetupRealClient()
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

	_, err = feedB.FollowersWithLimitAndSkip(20, 0)
	if err != nil {
		t.Fail()
	}

	testCleanUpFollows(client, []*FlatFeed{feedA, feedB, feedC})

}

func TestFlatFeedUnFollow(t *testing.T) {

	client, err := testSetupRealClient()
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

	err = feedA.Unfollow(feedB)
	if err != nil {
		t.Fail()
	}

	testCleanUpFollows(client, []*FlatFeed{feedA, feedB})

}

func TestFlatFeedUnFollowKeepingHistory(t *testing.T) {

	client, err := testSetupRealClient()
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

	err = feedA.UnfollowKeepingHistory(feedB)
	if err != nil {
		t.Fail()
	}

	testCleanUpFollows(client, []*FlatFeed{feedA, feedB})

}
