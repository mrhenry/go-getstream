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

	_, err = client.signer.GenerateFeedScopeToken(FeedContext, ReadAction, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(ActivitiesContext, WriteAction, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(FollowerContext, DeleteAction, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(AllContexts, AllActions, feed)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateFeedScopeToken(FeedContext, ReadAction, nil)
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

	_, err = client.signer.GenerateUserScopeToken(FeedContext, ReadAction, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(ActivitiesContext, WriteAction, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(FollowerContext, DeleteAction, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(AllContexts, AllActions, user)
	if err != nil {
		t.Fail()
	}

	_, err = client.signer.GenerateUserScopeToken(FeedContext, ReadAction, "")
	if err != nil {
		t.Fail()
	}
}
