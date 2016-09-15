package getstream_test

import (
	"fmt"
	"testing"
	"github.com/GetStream/stream-go"
)

func TestGeneralFeedBasic(t *testing.T) {

	client, err := getstream.New("a key", "a secret", "11111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	general := getstream.GeneralFeed{
		Client:   client,
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

	general.SignFeed(general.Client.Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Token() {
		t.Fail()
		return
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.GenerateToken(general.Client.Signer) {
		t.Fail()
		return
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != general.Signature() {
		t.Fail()
		return
	}
}

func TestFlatFeedBasic(t *testing.T) {

	client, err := getstream.New("a key", "a secret", "11111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	flatFeed := getstream.FlatFeed{
		Client:   client,
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

	flatFeed.SignFeed(flatFeed.Client.Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Token() {
		t.Fail()
		return
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.GenerateToken(flatFeed.Client.Signer) {
		t.Fail()
		return
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != flatFeed.Signature() {
		t.Fail()
		return
	}
}

func TestNotificationFeedBasic(t *testing.T) {

	client, err := getstream.New("a key", "a secret", "11111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	notificationFeed := getstream.NotificationFeed{
		Client:   client,
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

	notificationFeed.SignFeed(notificationFeed.Client.Signer)
	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Token() {
		t.Fail()
		return
	}

	if "NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.GenerateToken(notificationFeed.Client.Signer) {
		t.Fail()
		return
	}

	if "feedGroupfeedName NWH8lcFHfHYEc2xdMs2kOhM-oII" != notificationFeed.Signature() {
		t.Fail()
		return
	}
}
