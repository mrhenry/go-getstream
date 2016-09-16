package getstream_test

import (
	"fmt"
	"github.com/GetStream/stream-go"
	"testing"
)

func TestFeedSlug(t *testing.T) {
	slug, err := getstream.ValidateFeedSlug("foo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if slug != "foo" {
		fmt.Println("feedSlug not 'foo'")
		t.Fail()
	}

	slug, err = getstream.ValidateFeedSlug("f-o-o")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if slug != "f_o_o" {
		fmt.Println("feedSlug not 'f_o_o'")
		t.Fail()
	}
}

func TestFeedID(t *testing.T) {
	feedID, err := getstream.ValidateFeedID("123")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if feedID != "123" {
		fmt.Println("feedID not '123'")
		t.Fail()
	}

	feedID, err = getstream.ValidateFeedID("1-2-3")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if feedID != "1_2_3" {
		fmt.Println("feedSlug not '1_2_3'")
		t.Fail()
	}
}

func TestUserID(t *testing.T) {
	userID, err := getstream.ValidateUserID("123")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if userID != "123" {
		fmt.Println("userID not '123'")
		t.Fail()
	}

	userID, err = getstream.ValidateUserID("1-2-3")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if userID != "1_2_3" {
		fmt.Println("userSlug not '1_2_3'")
		t.Fail()
	}
}
