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

	client, err := New("a_key", "tfq2sdqpj9g446sbv653x3aqmgn33hsn8uzdc9jpskaw8mj6vsnhzswuwptuj9su", "123456", "us-east")
	if err != nil {
		t.Fail()
	}

	feed, err := client.FlatFeed("flat", "bob")
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(ScopeContextFeed, ScopeActionRead, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(ScopeContextActivities, ScopeActionWrite, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(ScopeContextFollower, ScopeActionDelete, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(ScopeContextAll, ScopeActionAll, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(ScopeContextFeed, ScopeActionRead, nil)
	if err != nil {
		t.Fail()
	}
}

func TestUserScopeToken(t *testing.T) {

	client, err := New("a_key", "tfq2sdqpj9g446sbv653x3aqmgn33hsn8uzdc9jpskaw8mj6vsnhzswuwptuj9su", "123456", "us-east")
	if err != nil {
		t.Fail()
	}

	user := "bob"

	_, err = client.signer.GenerateUserScopeToken(ScopeContextFeed, ScopeActionRead, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(ScopeContextActivities, ScopeActionWrite, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(ScopeContextFollower, ScopeActionDelete, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(ScopeContextAll, ScopeActionAll, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(ScopeContextFeed, ScopeActionRead, "")
	if err != nil {
		t.Fail()
	}
}
