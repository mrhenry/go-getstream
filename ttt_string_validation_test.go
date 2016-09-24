package getstream

import "testing"

func TestFeedSlug(t *testing.T) {
	feedSlug, err := ValidateFeedSlug("foo")
	if err != nil {
		t.Error(err)
	}
	if feedSlug != "foo" {
		t.Error("feedSlug not 'foo'")
	}

	feedSlug, err = ValidateFeedSlug("f-o-o")
	if err != nil {
		t.Error(err)
	}
	if feedSlug != "f_o_o" {
		t.Error("feedSlug not 'f_o_o'")
	}
}

func TestFeedID(t *testing.T) {
	feedID, err := ValidateFeedID("123")
	if err != nil {
		t.Error(err)
	}
	if feedID != "123" {
		t.Error("feedID not '123'")
	}

	feedID, err = ValidateFeedID("1-2-3")
	if err != nil {
		t.Error(err)
	}
	if feedID != "1_2_3" {
		t.Error("feedID not '1_2_3'")
	}
}

func TestUserID(t *testing.T) {
	userID, err := ValidateUserID("123")
	if err != nil {
		t.Error(err)
	}
	if userID != "123" {
		t.Error("userID not '123'")
	}

	userID, err = ValidateUserID("1-2-3")
	if err != nil {
		t.Error(err)
	}
	if userID != "1_2_3" {
		t.Error("userSlug not '1_2_3'")
	}
}
