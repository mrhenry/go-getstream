package getstream

import "testing"

func TestFlatFeedInputValidation(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		t.Fail()
	}

	_, err = client.FlatFeed("user", "099978b6-3b72-4f5c-bc43-247ba6ae2dd9")
	if err != nil {
		t.Fail()
	}

	_, err = client.FlatFeed("user", "tester@mail.com")
	if err == nil {
		t.Fail()
	}

}

func TestNotificationFeedInputValidation(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		t.Fail()
	}

	_, err = client.NotificationFeed("user", "099978b6-3b72-4f5c-bc43-247ba6ae2dd9")
	if err != nil {
		t.Fail()
	}

	_, err = client.NotificationFeed("user", "tester@mail.com")
	if err == nil {
		t.Fail()
	}

}

func TestClientInit(t *testing.T) {

	_, err := New("my_key", "my_secret", "111111", "!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		t.Fail()
	}

}

func TestClientBaseURL(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		t.Fail()
	}

	if "https://us-east-api.getstream.io/api/v1.0/" != client.BaseURL().String() {
		t.Fail()
	}
}

func TestClientAbsoluteURL(t *testing.T) {

	client, err := New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		t.Fail()
	}

	url, err := client.absoluteURL("user")
	if err != nil {
		t.Fail()
	}

	if "https://us-east-api.getstream.io/api/v1.0/user?api_key=my_key&location=us-east" != url.String() {
		t.Fail()
	}

}
