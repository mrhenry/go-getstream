package getstream_test

import (
	"os"

	getstream "github.com/GetStream/stream-go"
)

func PreTestSetup() (*getstream.Client, error) {
	return doTestSetup(&getstream.Config{
		APIKey:     os.Getenv("STREAM_API_KEY"),
		APISecret:  os.Getenv("STREAM_API_SECRET"),
		AppID:      os.Getenv("STREAM_APP_ID"),
		Location:   os.Getenv("STREAM_REGION"),
		TimeoutInt: 1000,
	})
}

func PreTestSetupWithToken() (*getstream.Client, error) {
	return doTestSetup(&getstream.Config{
		APIKey:     os.Getenv("STREAM_API_KEY"),
		Token:      os.Getenv("STREAM_API_SECRET"), // instead of APISecret
		AppID:      os.Getenv("STREAM_APP_ID"),
		Location:   os.Getenv("STREAM_REGION"),
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

func PostTestCleanUpFlatFeedFollows(client *getstream.Client, feeds []*getstream.FlatFeed) error {
	for _, feed := range feeds {
		followers, _ := feed.FollowersWithLimitAndSkip(300, 0)

		for _, follower := range followers {
			follower.Unfollow(client, feed)
		}
	}
	return nil
}
