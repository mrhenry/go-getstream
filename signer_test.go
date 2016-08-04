package getstream

import "testing"

func TestGenerateToken(t *testing.T) {

	signer := Signer{
		Secret: "test_secret",
	}
	token := signer.generateToken("some message")

	if token != "8SZVOYgCH6gy-ZjBTq_9vydr7TQ" {
		t.Fail()
	}
}

func TestURLSafe(t *testing.T) {

	signer := Signer{}

	result := signer.urlSafe("some+test/string=foo=")

	if result != "some-test_string=foo" {
		t.Fail()
	}
}
