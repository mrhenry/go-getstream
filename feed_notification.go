package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// NotificationFeed is a getstream NotificationFeed
// Use it to for CRUD on NotificationFeed Groups
type NotificationFeed struct {
	client   *Client
	FeedSlug string
	UserID   string
	token    string
}

// Client returns the Client associated with the NotificationFeed
func (f NotificationFeed) Client() *Client {
	return f.client
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *NotificationFeed) Signature() string {
	if f.Token() == "" {
		return f.feedIDWithoutColon()
	}
	return f.feedIDWithoutColon() + " " + f.Token()
}

// FeedID is the combo if the FeedSlug and UserID : "FeedSlug:UserID"
func (f *NotificationFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

func (f *NotificationFeed) feedIDWithoutColon() string {
	return f.FeedSlug + f.UserID
}

// SignFeed sets the token on a Feed
func (f *NotificationFeed) SignFeed(signer *Signer) {
	f.token = signer.generateToken(f.feedIDWithoutColon())
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

	resultBytes, err := f.Client().post(f, endpoint, f.Signature(), payload, nil)
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

	resultBytes, err := f.Client().post(f, endpoint, f.Signature(), payload, nil)
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

// MarkActivitiesAsRead marks activities as read for this feed
func (f *NotificationFeed) MarkActivitiesAsRead(activities []*NotificationFeedActivity) error {

	var ids []string
	for _, activity := range activities {
		ids = append(ids, activity.ID)
	}

	idStr := strings.Join(ids, ",")

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	_, err := f.Client().get(f, endpoint, f.Signature(), nil, map[string]string{
		"mark_read": idStr,
	})

	return err
}

// MarkActivitiesAsSeenWithLimit marks activities as seen for this feed
func (f *NotificationFeed) MarkActivitiesAsSeenWithLimit(limit int) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	_, err := f.Client().get(f, endpoint, f.Signature(), nil, map[string]string{
		"mark_seen": "true",
		"limit":     strconv.Itoa(limit),
	})

	return err
}

// Activities returns a list of Activities for a NotificationFeedGroup
func (f *NotificationFeed) Activities(input *GetNotificationFeedInput) (*GetNotificationFeedOutput, error) {

	var payload []byte
	var err error

	if input != nil {
		payload, err = json.Marshal(input)
		if err != nil {
			return nil, err
		}
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	result, err := f.Client().get(f, endpoint, f.Signature(), payload, nil)
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

	return f.Client().del(f, endpoint, f.Signature(), nil, nil)
}

// RemoveActivityByForeignID removes an Activity from a NotificationFeedGroup by ForeignID
func (f *NotificationFeed) RemoveActivityByForeignID(input *NotificationFeedActivity) error {

	if input.ForeignID == "" {
		return errors.New("no ForeignID")
	}

	r, err := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")
	if err != nil {
		return err
	}
	if !r.MatchString(input.ForeignID) {
		return errors.New("invalid ForeignID")
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + input.ForeignID + "/"

	return f.Client().del(f, endpoint, f.Signature(), nil, map[string]string{
		"foreign_id": "1",
	})
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

	_, err = f.Client().post(f, endpoint, f.Signature(), payload, nil)
	return err

}

// Unfollow is used to Unfollow a target Feed
func (f *NotificationFeed) Unfollow(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + target.FeedID().Value() + "/"

	return f.Client().del(f, endpoint, f.Signature(), nil, nil)

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

	return f.Client().del(f, endpoint, f.Signature(), payload, nil)

}

// FollowingWithLimitAndSkip returns a list of GeneralFeed followed by the current FlatFeed
func (f *NotificationFeed) FollowingWithLimitAndSkip(limit int, skip int) ([]*GeneralFeed, error) {

	var err error

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	payload, err := json.Marshal(&getNotificationFeedFollowersInput{
		Limit: limit,
		Skip:  skip,
	})
	if err != nil {
		return nil, err
	}

	resultBytes, err := f.Client().get(f, endpoint, f.Signature(), payload, nil)

	output := &getNotificationFeedFollowersOutput{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	var outputFeeds []*GeneralFeed
	for _, result := range output.Results {

		feed := GeneralFeed{}

		var match bool
		match, err = regexp.MatchString(`^.*?:.*?$`, result.FeedID)
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
