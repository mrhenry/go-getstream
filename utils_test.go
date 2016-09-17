package getstream

import "os"

func testSetup() (*Client, error) {

	testAPIKey := os.Getenv("key")
	testAPISecret := os.Getenv("secret")
	testAppID := os.Getenv("app_id")
	testRegion := os.Getenv("region")

	client := New(testAPIKey, testAPISecret, testAppID, testRegion)

	return client, nil

}

func testCleanUp(client *Client, flats []*Activity, notifications []*Activity, aggregations []*Activity) error {

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

func testCleanUpFollows(client *Client, flats []*FlatFeed) error {

	for _, flat := range flats {

		followers, _ := flat.FollowersWithLimitAndSkip(300, 0)

		for _, follower := range followers {
			follower.Unfollow(client, flat)
		}
	}
	return nil
}
