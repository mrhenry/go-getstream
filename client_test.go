package getstream_test

import (
	"os"
	"testing"

	getstream "github.com/GetStream/stream-go"
)

func TestFlatFeedInputValidation(t *testing.T) {
	client, err := getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.FlatFeed("user", "099978b6-3b72-4f5c-bc43-247ba6ae2dd9")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.FlatFeed("user", "tester@mail.com")
	if err == nil {
		t.Fatal(err)
	}
}

func TestNotificationFeedInputValidation(t *testing.T) {
	client, err := getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.NotificationFeed("user", "099978b6-3b72-4f5c-bc43-247ba6ae2dd9")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.NotificationFeed("user", "tester@mail.com")
	if err == nil {
		t.Fatal(err)
	}
}

func TestClientInit(t *testing.T) {
	_, err := getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  "!#@#$%ˆ&*((*=/*-+[]',.><"})
	if err == nil {
		t.Fatal(err)
	}

	_, err = getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  ""})
	if err != nil {
		t.Fatal(err)
	}

	_, err = getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientInitWithToken(t *testing.T) {
	serverClient, err := PreTestSetupWithToken()
	if err != nil {
		t.Fatal(err)
	}
	if serverClient.Signer == nil {
		t.Fatal("Required Signer is nil")
	}

	serverFeed, err := serverClient.FlatFeed("flat", "bob")
	if err != nil {
		t.Fatal(err)
	}

	token, err := serverClient.Signer.GenerateFeedScopeToken(
		getstream.ScopeContextAll,
		getstream.ScopeActionAll,
		serverFeed.FeedIDWithoutColon())
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("signer generated feed scope token is blank")
	}

	// now we're going to pass that token from above into a new client instead of APISecret

	clientClient, err := getstream.New(&getstream.Config{
		APIKey:   os.Getenv("key"),
		Token:    token, // pass token instead of API Secret
		AppID:    os.Getenv("app_id"),
		Location: os.Getenv("region")})
	if err != nil {
		t.Fatal(err)
	}

	feed, err := clientClient.FlatFeed("flat", "bob")
	if err != nil {
		t.Fatal(err)
	}

	activity, err := feed.AddActivity(&getstream.Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    getstream.FeedID("flat:eric"),
		Actor:     getstream.FeedID("flat:john"),
	})
	if err != nil {
		t.Fatal(err)
	}

	if activity.Verb != "post" && activity.ForeignID != "48d024fe-3752-467a-8489-23febd1dec4e" {
		t.Fail()
	}

	// tests passed, do cleanup

	err = PostTestCleanUp(clientClient, []*getstream.Activity{activity}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientBaseURL(t *testing.T) {
	client, err := getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	if string(client.BaseURL.String()) != "https://us-east-api.getstream.io/api/v1.0/" {
		t.Fatal()
	}
}

func TestClientAbsoluteURL(t *testing.T) {
	client, err := getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  "us-east"})
	if err != nil {
		t.Fatal(err)
	}

	url, err := client.AbsoluteURL("user")
	if err != nil {
		t.Fatal(err)
	}

	if url.String() != "https://us-east-api.getstream.io/api/v1.0/user?api_key=my_key&location=us-east" {
		t.Fatal(err)
	}

	client, err = getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  ""})
	if err != nil {
		t.Fatal(err)
	}

	url, err = client.AbsoluteURL("flat")
	if err != nil {
		t.Fatal(err)
	}

	if url.String() != "https://api.getstream.io/api/v1.0/flat?api_key=my_key&location=unspecified" {
		t.Fatal()
	}

	client, err = getstream.New(&getstream.Config{
		APIKey:    "my_key",
		APISecret: "my_secret",
		AppID:     "111111",
		Location:  ""})
	if err != nil {
		t.Fatal(err)
	}

	url, err = client.AbsoluteURL("!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		t.Fatal(err)
	}
}
