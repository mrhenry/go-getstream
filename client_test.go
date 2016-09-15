package getstream_test

import (
	"fmt"
	"os"
	"testing"

	getstream "github.com/GetStream/stream-go"
)

func TestFlatFeedInputValidation(t *testing.T) {

	client, err := getstream.New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = client.FlatFeed("user", "099978b6-3b72-4f5c-bc43-247ba6ae2dd9")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = client.FlatFeed("user", "tester@mail.com")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestNotificationFeedInputValidation(t *testing.T) {

	client, err := getstream.New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = client.NotificationFeed("user", "099978b6-3b72-4f5c-bc43-247ba6ae2dd9")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = client.NotificationFeed("user", "tester@mail.com")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestClientInit(t *testing.T) {

	_, err := getstream.New("my_key", "my_secret", "111111", "!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = getstream.New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = getstream.New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestClientInitWithToken(t *testing.T) {

	serverClient, err := getstream.PreTestSetup()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	serverFeed, err := serverClient.FlatFeed("flat", "bob")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	token, err := serverClient.Signer.GenerateFeedScopeToken(
		getstream.ScopeContextAll,
		getstream.ScopeActionAll,
		serverFeed.FeedIDWithoutColon())
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	testAPIKey := os.Getenv("key")
	testAppID := os.Getenv("app_id")
	testRegion := os.Getenv("region")

	clientClient, err := getstream.NewWithToken(testAPIKey, token, testAppID, testRegion)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	feed, err := clientClient.FlatFeed("flat", "bob")
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
		return
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
		return
	}

	err = getstream.PostTestCleanUp(serverClient, []*getstream.Activity{activity}, nil, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestClientBaseURL(t *testing.T) {

	client, err := getstream.New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if client.BaseURL.String() != "https://us-east-api.getstream.io/api/v1.0/"{
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestClientAbsoluteURL(t *testing.T) {

	client, err := getstream.New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	url, err := client.AbsoluteURL("user")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if "https://us-east-api.getstream.io/api/v1.0/user?api_key=my_key&location=us-east" != url.String() {
		fmt.Println(err)
		t.Fail()
		return
	}

	client, err = getstream.New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	url, err = client.AbsoluteURL("flat")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if "https://api.getstream.io/api/v1.0/flat?api_key=my_key&location=unspecified" != url.String() {
		fmt.Println(err)
		t.Fail()
		return
	}

	client, err = getstream.New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	url, err = client.AbsoluteURL("!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}
