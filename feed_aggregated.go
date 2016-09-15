package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// AggregatedFeed is a getstream AggregatedFeed
// Use it to for CRUD on AggregatedFeed Groups
type AggregatedFeed struct {
	client   *Client
	FeedSlug string
	UserID   string
	token    string
}

// Client returns the Client associated with the AggregatedFeed
func (f AggregatedFeed) Client() *Client {
	return f.client
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *AggregatedFeed) Signature() string {
	if f.Token() == "" {
		return f.feedIDWithoutColon()
	}
	return f.feedIDWithoutColon() + " " + f.Token()
}

// FeedID is the combo if the FeedSlug and UserID : "FeedSlug:UserID"
func (f *AggregatedFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

func (f *AggregatedFeed) feedIDWithoutColon() string {
	return f.FeedSlug + f.UserID
}

// Token returns the token of a Feed
func (f *AggregatedFeed) Token() string {
	return f.token
}

// GenerateToken returns a new Token for a Feed without setting it to the Feed
func (f *AggregatedFeed) GenerateToken(signer *Signer) string {
	if f.Client().Signer != nil {
		return signer.generateToken(f.FeedSlug + f.UserID)
	}
	return ""
}

// AddActivity is used to add an Activity to a AggregatedFeed
func (f *AggregatedFeed) AddActivity(activity *Activity) (*Activity, error) {

	payload, err := json.Marshal(activity)
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.Client().post(f, endpoint, payload, nil)
	if err != nil {
		return nil, err
	}

	output := &Activity{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	return output, err
}

// AddActivities is used to add multiple Activities to a NotificationFeed
func (f *AggregatedFeed) AddActivities(activities []*Activity) ([]*Activity, error) {

	payload, err := json.Marshal(map[string][]*Activity{
		"activities": activities,
	})
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.Client().post(f, endpoint, payload, nil)
	if err != nil {
		return nil, err
	}

	output := &postAggregatedFeedOutputActivities{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	return output.Activities, err
}

// Activities returns a list of Activities for a NotificationFeedGroup
func (f *AggregatedFeed) Activities(input *GetAggregatedFeedInput) (*GetAggregatedFeedOutput, error) {

	var payload []byte
	var err error

	if input != nil {
		payload, err = json.Marshal(input)
		if err != nil {
			return nil, err
		}
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	result, err := f.Client().get(f, endpoint, payload, nil)
	if err != nil {
		return nil, err
	}

	output := &getAggregatedFeedOutput{}
	err = json.Unmarshal(result, output)
	if err != nil {
		return nil, err
	}

	return output.output(), err
}

// RemoveActivity removes an Activity from a NotificationFeedGroup
func (f *AggregatedFeed) RemoveActivity(input *Activity) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + input.ID + "/"

	return f.Client().del(f, endpoint, nil, nil)
}

// RemoveActivityByForeignID removes an Activity from a NotificationFeedGroup by ForeignID
func (f *AggregatedFeed) RemoveActivityByForeignID(input *Activity) error {

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

	return f.Client().del(f, endpoint, nil, map[string]string{
		"foreign_id": "1",
	})
}

// FollowFeedWithCopyLimit sets a Feed to follow another target Feed
// CopyLimit is the maximum number of Activities to Copy from History
func (f *AggregatedFeed) FollowFeedWithCopyLimit(target *FlatFeed, copyLimit int) error {
	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	input := postAggregatedFeedFollowingInput{
		Target:            target.FeedID().Value(),
		ActivityCopyLimit: copyLimit,
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}

	_, err = f.Client().post(f, endpoint, payload, nil)
	return err

}

// Unfollow is used to Unfollow a target Feed
func (f *AggregatedFeed) Unfollow(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + target.FeedID().Value() + "/"

	return f.Client().del(f, endpoint, nil, nil)

}

// UnfollowKeepingHistory is used to Unfollow a target Feed while keeping the History
// this means that Activities already visibile will remain
func (f *AggregatedFeed) UnfollowKeepingHistory(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + target.FeedID().Value() + "/"

	payload, err := json.Marshal(map[string]string{
		"keep_history": "1",
	})
	if err != nil {
		return err
	}

	return f.Client().del(f, endpoint, payload, nil)

}

// FollowingWithLimitAndSkip returns a list of GeneralFeed followed by the current FlatFeed
func (f *AggregatedFeed) FollowingWithLimitAndSkip(limit int, skip int) ([]*GeneralFeed, error) {

	var err error

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	payload, err := json.Marshal(&getAggregatedFeedFollowersInput{
		Limit: limit,
		Skip:  skip,
	})
	if err != nil {
		return nil, err
	}

	resultBytes, err := f.Client().get(f, endpoint, payload, nil)

	output := &getAggregatedFeedFollowersOutput{}
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
