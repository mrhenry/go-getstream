package getstream

import "testing"

func TestGenerateToken(t *testing.T) {

	signer := Signer{
		Secret: "test_secret",
	}
	token := signer.generateToken("some message")

	if token != "8SZVOYgCH6gy-ZjBTq_9vydr7TQ" {
		t.Fail()
		return
	}
}

func TestURLSafe(t *testing.T) {

	signer := Signer{}

	result := signer.urlSafe("some+test/string=foo=")

	if result != "some-test_string=foo" {
		t.Fail()
		return
	}
}

func TestFeedScopeToken(t *testing.T) {

	opts := ServerOptions("a_key", "tfq2sdqpj9g446sbv653x3aqmgn33hsn8uzdc9jpskaw8mj6vsnhzswuwptuj9su", "123456", "us-east")
	client, err := New(opts)
	if err != nil {
		t.Fail()
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(ScopeContextFeed, ScopeActionRead, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(ScopeContextActivities, ScopeActionWrite, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(ScopeContextFollower, ScopeActionDelete, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(ScopeContextAll, ScopeActionAll, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.Signer.GenerateFeedScopeToken(ScopeContextFeed, ScopeActionRead, nil)
	if err != nil {
		t.Fail()
	}
}

func TestUserScopeToken(t *testing.T) {

	opts := ServerOptions("a_key", "tfq2sdqpj9g446sbv653x3aqmgn33hsn8uzdc9jpskaw8mj6vsnhzswuwptuj9su", "123456", "us-east")
	client, err := New(opts)
	if err != nil {
		t.Fail()
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
