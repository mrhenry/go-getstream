package getstream

import (
	"fmt"
	"testing"
)

func TestGeneralFeedBasic(t *testing.T) {

	opts := ServerOptions("a key", "a secret", "11111", "us-east")
	client, err := New(opts)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	general := GeneralFeed{
		client:   client,
		FeedSlug: "feedGroup",
		UserID:   "feedName",
	}

	if "feedGroupfeedName" != general.Signature() {
		t.Fail()
		return
	}

	if "feedGroup:feedName" != string(general.FeedID()) {
		t.Fail()
		return
	}

	general.SignFeed(general.Client().Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Token() {
		t.Fail()
		return
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.GenerateToken(general.Client().Signer) {
		t.Fail()
		return
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Signature() {
		t.Fail()
		return
	}
}

func TestFlatFeedBasic(t *testing.T) {

	opts := ServerOptions("a key", "a secret", "11111", "us-east")
	client, err := New(opts)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	flatFeed := FlatFeed{
		client:   client,
		FeedSlug: "feedGroup",
		UserID:   "feedName",
	}

	if "feedGroupfeedName" != flatFeed.Signature() {
		t.Fail()
		return
	}

	if "feedGroup:feedName" != string(flatFeed.FeedID()) {
		t.Fail()
		return
	}

	flatFeed.SignFeed(flatFeed.Client().Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Token() {
		t.Fail()
		return
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.GenerateToken(flatFeed.Client().Signer) {
		t.Fail()
		return
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Signature() {
		t.Fail()
		return
	}
}

func TestNotificationFeedBasic(t *testing.T) {

	opts := ServerOptions("a key", "a secret", "11111", "us-east")
	client, err := New(opts)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	notificationFeed := NotificationFeed{
		client:   client,
		FeedSlug: "feedGroup",
		UserID:   "feedName",
	}

	if "feedGroupfeedName" != notificationFeed.Signature() {
		t.Fail()
		return
	}

	if "feedGroup:feedName" != string(notificationFeed.FeedID()) {
		t.Fail()
		return
	}

	notificationFeed.SignFeed(notificationFeed.Client().Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Token() {
		t.Fail()
		return
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.GenerateToken(notificationFeed.Client().Signer) {
		t.Fail()
		return
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Signature() {
		t.Fail()
		return
	}
}
