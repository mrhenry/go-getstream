package getstream

import (
	"fmt"
	"os"
	"testing"
)

func TestFlatFeedInputValidation(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
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

	client, err := New("my_key", "my_secret", "111111", "us-east")
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

	_, err := New("my_key", "my_secret", "111111", "!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestClientInitWithToken(t *testing.T) {

	serverClient, err := testSetup()
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

	token, err := serverClient.Signer.GenerateFeedScopeToken(ScopeContextAll, ScopeActionAll, serverFeed)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	testAPIKey := os.Getenv("key")
	testAppID := os.Getenv("app_id")
	testRegion := os.Getenv("region")

	clientClient, err := NewWithToken(testAPIKey, token, testAppID, testRegion)
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

	activity, err := feed.AddActivity(&Activity{
		Verb:      "post",
		ForeignID: "48d024fe-3752-467a-8489-23febd1dec4e",
		Object:    FeedID("flat:eric"),
		Actor:     FeedID("flat:john"),
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

	err = testCleanUp(serverClient, []*Activity{activity}, nil, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestClientBaseURL(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if "https://us-east-api.getstream.io/api/v1.0/" != client.baseURL.String() {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func TestClientAbsoluteURL(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	url, err := client.absoluteURL("user")
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

	client, err = New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	url, err = client.absoluteURL("flat")
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

	client, err = New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	url, err = client.absoluteURL("!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}

func TestClientRequestFails(t *testing.T) {

	client, err := testSetup()
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

	_, err = client.request(feed, "GET", "!@#$%", nil, nil)
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	if err.Error() != string(`parse !@#$%: invalid URL escape "%"`) {
		fmt.Println(err)
		t.Fail()
		return
	}

	_, err = client.request(feed, "!@#$%", "somepath", nil, nil)
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	if err.Error() != string(`net/http: invalid method "!@#$%"`) {
		fmt.Println(err)
		t.Fail()
		return
	}

}
