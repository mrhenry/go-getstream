package getstream

import "testing"

func TestGeneralFeedBasic(t *testing.T) {
	client, err := New(&Config{
		APIKey:    "a key",
		APISecret: "a secret",
		AppID:     "11111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	general := GeneralFeed{
		Client:   client,
		feedSlug: "feedGroup",
		userID:   "feedName",
	}

	if "feedGroupfeedName" != general.Signature() {
		t.Fatal()
	}

	if "feedGroup:feedName" != string(general.FeedIDWithColon()) {
		t.Fatal()
	}

	general.SignFeed(general.Client.Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Token() {
		t.Fatal()
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.GenerateToken(general.Client.Signer) {
		t.Fatal()
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Signature() {
		t.Fatal()
	}
}

func TestFlatFeedBasic(t *testing.T) {
	client, err := New(&Config{
		APIKey:    "a key",
		APISecret: "a secret",
		AppID:     "11111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	flatFeed := FlatFeed{
		Client:   client,
		feedSlug: "feedGroup",
		userID:   "feedName",
	}

	if "feedGroupfeedName" != flatFeed.Signature() {
		t.Fatal()
	}

	if "feedGroup:feedName" != string(flatFeed.FeedIDWithColon()) {
		t.Fatal()
	}

	flatFeed.SignFeed(flatFeed.Client.Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Token() {
		t.Fatal()
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.GenerateToken(flatFeed.Client.Signer) {
		t.Fatal()
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Signature() {
		t.Fatal()
	}
}

func TestNotificationFeedBasic(t *testing.T) {
	client, err := New(&Config{
		APIKey:    "a key",
		APISecret: "a secret",
		AppID:     "11111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	notificationFeed := NotificationFeed{
		Client:   client,
		feedSlug: "feedGroup",
		userID:   "feedName",
	}

	if "feedGroupfeedName" != notificationFeed.Signature() {
		t.Fatal()
	}

	if "feedGroup:feedName" != string(notificationFeed.FeedIDWithColon()) {
		t.Fatal()
	}

	notificationFeed.SignFeed(notificationFeed.Client.Signer)

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Token() {
		t.Fatal()
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.GenerateToken(notificationFeed.Client.Signer) {
		t.Fatal()
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Signature() {
		t.Fatal()
	}
}
