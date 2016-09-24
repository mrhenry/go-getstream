package getstream

import "testing"

func TestGenerateToken(t *testing.T) {

	signer := Signer{
		Secret: "test_secret",
	}

	token := signer.GenerateToken("some message")
	if token != "8SZVOYgCH6gy-ZjBTq_9vydr7TQ" {
		t.Fail()
		return
	}
}

func TestURLSafe(t *testing.T) {

	signer := Signer{}

	result := signer.URLSafe("some+test/string=foo=")
	if result != "some-test_string=foo" {
		t.Fail()
		return
	}
}

func TestFeedScopeToken(t *testing.T) {

	client, err := New(&Config{
		APIKey:    "a_key",
		APISecret: "tfq2sdqpj9g446sbv653x3aqmgn33hsn8uzdc9jpskaw8mj6vsnhzswuwptuj9su",
		AppID:     "123456",
		Location:  "us-east"})
	if err != nil {
		t.Fail()
		return
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(
		ScopeContextFeed,
		ScopeActionRead,
		feed.FeedIDWithoutColon())
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(
		ScopeContextActivities,
		ScopeActionWrite,
		feed.FeedIDWithoutColon())
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(
		ScopeContextFollower,
		ScopeActionDelete,
		feed.FeedIDWithoutColon())
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(
		ScopeContextAll,
		ScopeActionAll,
		feed.FeedIDWithoutColon())
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(ScopeContextFeed, ScopeActionRead, "")
	if err != nil {
		t.Fail()
	}
}

func TestUserScopeToken(t *testing.T) {

	client, err := New(&Config{
		APIKey:    "a_key",
		APISecret: "tfq2sdqpj9g446sbv653x3aqmgn33hsn8uzdc9jpskaw8mj6vsnhzswuwptuj9su",
		AppID:     "123456",
		Location:  "us-east"})
	if err != nil {
		t.Fail()
		return
	}

	user := "bob"

	_, err = client.Signer.GenerateUserScopeToken(ScopeContextFeed, ScopeActionRead, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateUserScopeToken(ScopeContextActivities, ScopeActionWrite, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateUserScopeToken(ScopeContextFollower, ScopeActionDelete, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateUserScopeToken(ScopeContextAll, ScopeActionAll, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateUserScopeToken(ScopeContextFeed, ScopeActionRead, "")
	if err != nil {
		t.Fail()
	}
}
