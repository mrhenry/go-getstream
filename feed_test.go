package getstream

import "testing"

func TestGeneralFeedBasic(t *testing.T) {

	client, err := New("a key", "a secret", "11111", "us-east")
	if err != nil {
		t.Fail()
	}

	general := GeneralFeed{
		Client:   client,
		FeedSlug: "feedGroup",
		UserID:   "feedName",
	}

	if "feedGroupfeedName" != general.Signature() {
		t.Fail()
	}

	if "feedGroup:feedName" != string(general.FeedID()) {
		t.Fail()
	}

	general.SignFeed(general.Client.signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Token() {
		t.Fail()
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.GenerateToken(general.Client.signer) {
		t.Fail()
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Signature() {
		t.Fail()
	}
}

func TestFlatFeedBasic(t *testing.T) {

	client, err := New("a key", "a secret", "11111", "us-east")
	if err != nil {
		t.Fail()
	}

	flatFeed := FlatFeed{
		Client:   client,
		FeedSlug: "feedGroup",
		UserID:   "feedName",
	}

	if "feedGroupfeedName" != flatFeed.Signature() {
		t.Fail()
	}

	if "feedGroup:feedName" != string(flatFeed.FeedID()) {
		t.Fail()
	}

	flatFeed.SignFeed(flatFeed.Client.signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Token() {
		t.Fail()
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.GenerateToken(flatFeed.Client.signer) {
		t.Fail()
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Signature() {
		t.Fail()
	}
}

func TestNotificationFeedBasic(t *testing.T) {

	client, err := New("a key", "a secret", "11111", "us-east")
	if err != nil {
		t.Fail()
	}

	notificationFeed := NotificationFeed{
		Client:   client,
		FeedSlug: "feedGroup",
		UserID:   "feedName",
	}

	if "feedGroupfeedName" != notificationFeed.Signature() {
		t.Fail()
	}

	if "feedGroup:feedName" != string(notificationFeed.FeedID()) {
		t.Fail()
	}

	notificationFeed.SignFeed(notificationFeed.Client.signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Token() {
		t.Fail()
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.GenerateToken(notificationFeed.Client.signer) {
		t.Fail()
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Signature() {
		t.Fail()
	}
}
