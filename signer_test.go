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

func TestScopeToken(t *testing.T) {

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

	_, err = client.signer.GenerateFeedScopeToken(FeedContext, ReadAction, nil)
	if err != nil {
		t.Fail()
	}
}
