package getstream_test

import (
	"fmt"
	getstream "github.com/GetStream/stream-go"
	"os"
)

func PreTestSetup() (*getstream.Client, error) {
	return doTestSetup(&getstream.Config{
		APIKey:     os.Getenv("key"),
		APISecret:  os.Getenv("secret"),
		AppID:      os.Getenv("app_id"),
		Location:   os.Getenv("region"),
		TimeoutInt: 1000,
	})
}

func PreTestSetupWithToken() (*getstream.Client, error) {
	return doTestSetup(&getstream.Config{
		APIKey:     os.Getenv("key"),
		Token:      os.Getenv("secret"), // instead of APISecret
		AppID:      os.Getenv("app_id"),
		Location:   os.Getenv("region"),
		TimeoutInt: 1000,
	})
}

func doTestSetup(cfg *getstream.Config) (*getstream.Client, error) {
	return getstream.New(cfg)
}

func PostTestCleanUp(
	client *getstream.Client,
	flats []*getstream.Activity,
	notifications []*getstream.Activity,
	aggregations []*getstream.Activity) error {

	fmt.Println("Cleanup, aisle 1")

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

	if len(aggregations) > 0 {

		feed, err := client.AggregatedFeed("aggregated", "bob")
		if err != nil {
			return err
		}

		for _, activity := range aggregations {
			err := feed.RemoveActivity(activity)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func PostTestCleanUpFollows(client *getstream.Client, flats []*getstream.FlatFeed) error {
	fmt.Println("Cleanup, aisle 2")
	for _, flat := range flats {

		followers, _ := flat.FollowersWithLimitAndSkip(300, 0)

		for _, follower := range followers {
			follower.Unfollow(client, flat)
		}
	}
	return nil
}
