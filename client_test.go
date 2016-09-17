package getstream

import (
	"fmt"
	"testing"
)

func TestClientInit(t *testing.T) {

	_, err := New("my_key", "my_secret", "111111", "!#@#$%ˆ&*((*=/*-+[]',.><")
	if err == nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	client, err := New("my_key", "my_secret", "111111", "")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if client.Key != "my_key" {
		t.Fail()
		return
	}
	if client.Secret != "my_secret" {
		t.Fail()
		return
	}
	if client.AppID != "111111" {
		t.Fail()
		return
	}
	if client.Location != "" {
		t.Fail()
		return
	}
	if client.baseURL.String() != "https://api.getstream.io/api/v1.0/" {
		fmt.Println(client.baseURL.String())
		t.Fail()
		return
	}
	if client.Signer == nil {
		t.Fail()
		return
	}
	if client.HTTP == nil {
		t.Fail()
		return
	}

	client, err = New("my_key", "my_secret", "111111", "us-east")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if client.Key != "my_key" {
		t.Fail()
		return
	}
	if client.Secret != "my_secret" {
		t.Fail()
		return
	}
	if client.AppID != "111111" {
		t.Fail()
		return
	}
	if client.Location != "us-east" {
		t.Fail()
		return
	}
	if client.baseURL.String() != "https://us-east-api.getstream.io/api/v1.0/" {
		fmt.Println(client.baseURL.String())
		t.Fail()
		return
	}
	if client.Signer == nil {
		t.Fail()
		return
	}
	if client.HTTP == nil {
		t.Fail()
		return
	}

}

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
