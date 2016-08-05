package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"
)

// NotificationFeedActivity is a getstream Activity
// Use it to post activities to NotificationFeeds
// It is also the response from NotificationFeed Fetch and List Requests
type NotificationFeedActivity struct {
	ID        string
	Actor     FeedID
	Verb      string
	Object    FeedID
	Target    FeedID
	Origin    FeedID
	TimeStamp *time.Time

	ForeignID string
	Data      json.RawMessage

	To []Feed
}

func (a NotificationFeedActivity) input() (*postNotificationFeedInputActivity, error) {

	input := postNotificationFeedInputActivity{
		ID:     a.ID,
		Actor:  string(a.Actor),
		Verb:   a.Verb,
		Object: string(a.Object),
		Target: string(a.Target),
		Data:   a.Data,
	}

	if a.ForeignID != "" {
		r, err := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")
		if err != nil {
			return nil, err
		}
		if !r.MatchString(a.ForeignID) {
			return nil, errors.New("invalid ForeignID")
		}

		input.ForeignID = a.ForeignID
	}

	input.To = []string{}

	if a.TimeStamp == nil {
		input.RawTime = time.Now().Format("2006-01-02T15:04:05.999999")
	} else {
		input.RawTime = a.TimeStamp.Format("2006-01-02T15:04:05.999999")
	}

	for _, feed := range a.To {
		to := string(feed.FeedID())
		if feed.Token() != "" {
			to += " " + feed.Token()
		}
		input.To = append(input.To, to)
	}

	return &input, nil
}

type postNotificationFeedInputActivity struct {
	ID        string          `json:"id,omitempty"`
	Actor     string          `json:"actor"`
	Verb      string          `json:"verb"`
	Object    string          `json:"object"`
	Target    string          `json:"target,omitempty"`
	RawTime   string          `json:"time,omitempty"`
	ForeignID string          `json:"foreign_id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
	To        []string        `json:"to,omitempty"`
}

type postNotificationFeedOutputActivity struct {
	ID        string          `json:"id,omitempty"`
	Actor     string          `json:"actor"`
	Verb      string          `json:"verb"`
	Object    string          `json:"object"`
	Target    string          `json:"target,omitempty"`
	RawTime   string          `json:"time,omitempty"`
	ForeignID string          `json:"foreign_id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
	To        [][]string      `json:"to,omitempty"`
}

type postNotificationFeedOutputActivities struct {
	Activities []*postNotificationFeedOutputActivity `json:"activities"`
}

func (a postNotificationFeedOutputActivity) activity() *NotificationFeedActivity {

	activity := NotificationFeedActivity{
		ID:        a.ID,
		Actor:     FeedID(a.Actor),
		Verb:      a.Verb,
		Object:    FeedID(a.Object),
		Target:    FeedID(a.Target),
		ForeignID: a.ForeignID,
		Data:      a.Data,
	}

	if a.RawTime != "" {
		timeStamp, err := time.Parse("2006-01-02T15:04:05.999999", a.RawTime)
		if err == nil {
			activity.TimeStamp = &timeStamp
		}
	}

	for _, slice := range a.To {
		parseNotificationFeedToParams(slice, &activity)
	}
	return &activity
}

// GetNotificationFeedInput is used to Get a list of Activities from a NotificationFeed
type GetNotificationFeedInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE int `json:"id_gte,omitempty"`
	IDGT  int `json:"id_gt,omitempty"`
	IDLTE int `json:"id_lte,omitempty"`
	IDLT  int `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}

// GetNotificationFeedOutput is the response from a NotificationFeed Activities Get Request
type GetNotificationFeedOutput struct {
	Duration string
	Next     string
	Results  []*struct {
		Activities    []*NotificationFeedActivity
		ActivityCount int
		ActorCount    int
		CreatedAt     string
		Group         string
		ID            string
		IsRead        bool
		IsSeen        bool
		UpdatedAt     string
		Verb          string
	}
	Unread int
	Unseen int
}

type getNotificationFeedOutput struct {
	Duration string                             `json:"duration"`
	Next     string                             `json:"next"`
	Results  []*getNotificationFeedOutputResult `json:"results"`
	Unread   int                                `json:"unread"`
	Unseen   int                                `json:"unseen"`
}

func (a getNotificationFeedOutput) output() *GetNotificationFeedOutput {

	output := GetNotificationFeedOutput{
		Duration: a.Duration,
		Next:     a.Next,
		Unread:   a.Unread,
		Unseen:   a.Unseen,
	}

	var results []*struct {
		Activities    []*NotificationFeedActivity
		ActivityCount int
		ActorCount    int
		CreatedAt     string
		Group         string
		ID            string
		IsRead        bool
		IsSeen        bool
		UpdatedAt     string
		Verb          string
	}

	for _, result := range a.Results {

		outputResult := struct {
			Activities    []*NotificationFeedActivity
			ActivityCount int
			ActorCount    int
			CreatedAt     string
			Group         string
			ID            string
			IsRead        bool
			IsSeen        bool
			UpdatedAt     string
			Verb          string
		}{
			ActivityCount: result.ActivityCount,
			ActorCount:    result.ActorCount,
			CreatedAt:     result.CreatedAt,
			Group:         result.Group,
			ID:            result.ID,
			IsRead:        result.IsRead,
			IsSeen:        result.IsSeen,
			UpdatedAt:     result.UpdatedAt,
			Verb:          result.Verb,
		}

		for _, activity := range result.Activities {
			outputResult.Activities = append(outputResult.Activities, activity.activity())
		}

		results = append(results, &outputResult)
	}

	output.Results = results

	return &output
}

type getNotificationFeedOutputResult struct {
	Activities    []*getNotificationFeedOutputActivity `json:"activities"`
	ActivityCount int                                  `json:"activity_count"`
	ActorCount    int                                  `json:"actor_count"`
	CreatedAt     string                               `json:"created_at"`
	Group         string                               `json:"group"`
	ID            string                               `json:"id"`
	IsRead        bool                                 `json:"is_read"`
	IsSeen        bool                                 `json:"is_seen"`
	UpdatedAt     string                               `json:"updated_at"`
	Verb          string                               `json:"verb"`
}

type getNotificationFeedOutputActivity struct {
	ID        string          `json:"id,omitempty"`
	Actor     string          `json:"actor"`
	Verb      string          `json:"verb"`
	Object    string          `json:"object"`
	Target    string          `json:"target,omitempty"`
	RawTime   string          `json:"time,omitempty"`
	To        []string        `json:"to,omitempty"`
	Origin    string          `json:"origin"`
	ForeignID string          `json:"foreign_id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

func (a getNotificationFeedOutputActivity) activity() *NotificationFeedActivity {

	activity := NotificationFeedActivity{
		ID:        a.ID,
		Actor:     FeedID(a.Actor),
		Verb:      a.Verb,
		Object:    FeedID(a.Object),
		Target:    FeedID(a.Target),
		Origin:    FeedID(a.Origin),
		ForeignID: a.ForeignID,
		Data:      a.Data,
	}

	if a.RawTime != "" {
		timeStamp, err := time.Parse("2006-01-02T15:04:05.999999", a.RawTime)
		if err == nil {
			activity.TimeStamp = &timeStamp
		}
	}

	parseNotificationFeedToParams(a.To, &activity)

	return &activity
}

func parseNotificationFeedToParams(to []string, activity *NotificationFeedActivity) {

	for _, to := range to {

		feed := GeneralFeed{}

		match, err := regexp.MatchString(`^.*?:.*? .*?$`, to)
		if err != nil {
			continue
		}

		if match {
			firstSplit := strings.Split(to, ":")
			secondSplit := strings.Split(firstSplit[1], " ")

			feed.FeedSlug = firstSplit[0]
			feed.UserID = secondSplit[0]
			feed.token = secondSplit[1]
		}

		activity.To = append(activity.To, &feed)
	}

}

type getNotificationFeedFollowersInput struct {
	Limit int `json:"limit"`
	Skip  int `json:"offset"`
}

type getNotificationFeedFollowersOutput struct {
	Duration string                                      `json:"duration"`
	Results  []*getNotificationFeedFollowersOutputResult `json:"results"`
}

type getNotificationFeedFollowersOutputResult struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	FeedID    string `json:"feed_id"`
	TargetID  string `json:"target_id"`
}

type postNotificationFeedFollowingInput struct {
	Target            string `json:"target"`
	ActivityCopyLimit int    `json:"activity_copy_limit"`
}
