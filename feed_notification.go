package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// NotificationFeed is a getstream NotificationFeed
// Use it to for CRUD on NotificationFeed Groups
type NotificationFeed struct {
	Client   *Client
	FeedSlug string
	UserID   string
	token    string
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *NotificationFeed) Signature() string {
	if f.Token() == "" {
		return f.FeedSlug + f.UserID
	}
	return f.FeedSlug + f.UserID + " " + f.Token()
}

// FeedID is the combo if the FeedSlug and UserID : "FeedSlug:UserID"
func (f *NotificationFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

// SignFeed sets the token on a Feed
func (f *NotificationFeed) SignFeed(signer *Signer) {
	f.token = signer.generateToken(f.FeedSlug + f.UserID)
}

// Token returns the token of a Feed
func (f *NotificationFeed) Token() string {
	return f.token
}

// GenerateToken returns a new Token for a Feed without setting it to the Feed
func (f *NotificationFeed) GenerateToken(signer *Signer) string {
	return signer.generateToken(f.FeedSlug + f.UserID)
}

// AddActivity is used to add an Activity to a NotificationFeed
func (f *NotificationFeed) AddActivity(activity *NotificationFeedActivity) (*NotificationFeedActivity, error) {

	payload, err := json.Marshal(activity)
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.post(endpoint, f.Signature(), payload)
	if err != nil {
		return nil, err
	}

	output := &NotificationFeedActivity{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	return output, err
}

// AddActivities is used to add multiple Activities to a NotificationFeed
func (f *NotificationFeed) AddActivities(activities []*NotificationFeedActivity) ([]*NotificationFeedActivity, error) {

	payload, err := json.Marshal(map[string][]*NotificationFeedActivity{
		"activities": activities,
	})
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.post(endpoint, f.Signature(), payload)
	if err != nil {
		return nil, err
	}

	output := &postNotificationFeedOutputActivities{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	return output.Activities, err
}

// Activities returns a list of Activities for a NotificationFeedGroup
func (f *NotificationFeed) Activities(input *GetNotificationFeedInput) (*GetNotificationFeedOutput, error) {

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	result, err := f.get(endpoint, f.Signature(), payload)
	if err != nil {
		return nil, err
	}

	output := &getNotificationFeedOutput{}
	err = json.Unmarshal(result, output)
	if err != nil {
		return nil, err
	}

	return output.output(), err
}

// RemoveActivity removes an Activity from a NotificationFeedGroup
func (f *NotificationFeed) RemoveActivity(input *NotificationFeedActivity) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + input.ID + "/"

	return f.del(endpoint, f.Signature(), nil)
}

// RemoveActivityByForeignID removes an Activity from a NotificationFeedGroup by ForeignID
func (f *NotificationFeed) RemoveActivityByForeignID(input *NotificationFeedActivity) error {

	r, err := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")
	if err != nil {
		return err
	}
	if !r.MatchString(input.ForeignID) {
		return errors.New("invalid ForeignID")
	}

	if input.ForeignID == "" {
		return errors.New("no ForeignID")
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + input.ForeignID + "/"

	payload, err := json.Marshal(map[string]string{
		"foreign_id": "1",
	})
	if err != nil {
		return err
	}

	return f.del(endpoint, f.Signature(), payload)
}

// FollowFeedWithCopyLimit sets a Feed to follow another target Feed
// CopyLimit is the maximum number of Activities to Copy from History
func (f *NotificationFeed) FollowFeedWithCopyLimit(target *FlatFeed, copyLimit int) error {
	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	input := postNotificationFeedFollowingInput{
		Target:            target.FeedID().Value(),
		ActivityCopyLimit: copyLimit,
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}

	_, err = f.post(endpoint, f.Signature(), payload)
	return err

}

// Unfollow is used to Unfollow a target Feed
func (f *NotificationFeed) Unfollow(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + target.FeedID().Value() + "/"

	return f.del(endpoint, f.Signature(), nil)

}

// UnfollowKeepingHistory is used to Unfollow a target Feed while keeping the History
// this means that Activities already visibile will remain
func (f *NotificationFeed) UnfollowKeepingHistory(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + target.FeedID().Value() + "/"

	payload, err := json.Marshal(map[string]string{
		"keep_history": "1",
	})
	if err != nil {
		return err
	}

	return f.del(endpoint, f.Signature(), payload)

}

// FollowingWithLimitAndSkip returns a list of GeneralFeed followed by the current FlatFeed
func (f *NotificationFeed) FollowingWithLimitAndSkip(limit int, skip int) ([]*GeneralFeed, error) {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	payload, err := json.Marshal(&getNotificationFeedFollowersInput{
		Limit: limit,
		Skip:  skip,
	})
	if err != nil {
		return nil, err
	}

	resultBytes, err := f.get(endpoint, f.Signature(), payload)

	output := &getNotificationFeedFollowersOutput{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	var outputFeeds []*GeneralFeed
	for _, result := range output.Results {

		feed := GeneralFeed{}

		match, err := regexp.MatchString(`^.*?:.*?$`, result.FeedID)
		if err != nil {
			continue
		}

		if match {
			firstSplit := strings.Split(result.TargetID, ":")

			feed.FeedSlug = firstSplit[0]
			feed.UserID = firstSplit[1]
		}

		outputFeeds = append(outputFeeds, &feed)
	}

	return outputFeeds, err

}
