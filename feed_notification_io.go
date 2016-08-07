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
	Data      *json.RawMessage
	MetaData  map[string]string

	To []Feed
}

func (a NotificationFeedActivity) MarshalJSON() ([]byte, error) {

	payload := make(map[string]interface{})

	for key, value := range a.MetaData {
		payload[key] = value
	}

	payload["actor"] = string(a.Actor)
	payload["verb"] = string(a.Verb)
	payload["object"] = string(a.Object)
	payload["origin"] = string(a.Origin)

	if a.ID != "" {
		payload["id"] = a.ID
	}
	if a.Target != "" {
		payload["target"] = a.Target
	}

	if a.Data != nil {
		payload["data"] = a.Data
	}

	if a.ForeignID != "" {
		r, err := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")
		if err != nil {
			return nil, err
		}
		if !r.MatchString(a.ForeignID) {
			return nil, errors.New("invalid ForeignID")
		}
		payload["foreign_id"] = a.ForeignID
	}

	if a.TimeStamp == nil {
		payload["time"] = time.Now().Format("2006-01-02T15:04:05.999999")
	} else {
		payload["time"] = a.TimeStamp.Format("2006-01-02T15:04:05.999999")
	}

	var tos []string
	for _, feed := range a.To {
		to := string(feed.FeedID())
		if feed.Token() != "" {
			to += " " + feed.Token()
		}
		tos = append(tos, to)
	}

	if len(tos) > 0 {
		payload["to"] = a.To
	}

	return json.Marshal(payload)

}

func (a *NotificationFeedActivity) UnmarshalJSON(b []byte) (err error) {

	rawPayload := make(map[string]*json.RawMessage)
	metadata := make(map[string]string)

	json.Unmarshal(b, &rawPayload)

	for key, value := range rawPayload {
		lowerKey := strings.ToLower(key)

		if value == nil {
			continue
		}

		if lowerKey == "id" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.ID = strValue
		} else if lowerKey == "actor" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.Actor = FeedID(strValue)
		} else if lowerKey == "verb" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.Verb = strValue
		} else if lowerKey == "foreign_id" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.ForeignID = strValue
		} else if lowerKey == "object" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.Object = FeedID(strValue)
		} else if lowerKey == "origin" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.Origin = FeedID(strValue)
		} else if lowerKey == "target" {
			var strValue string
			json.Unmarshal(*value, &strValue)
			a.Target = FeedID(strValue)
		} else if lowerKey == "time" {
			var strValue string
			err := json.Unmarshal(*value, &strValue)
			if err != nil {
				continue
			}
			timeStamp, err := time.Parse("2006-01-02T15:04:05.999999", strValue)
			if err != nil {
				continue
			}
			a.TimeStamp = &timeStamp
		} else if lowerKey == "data" {
			a.Data = value
		} else if lowerKey == "to" {

			var to1D []string
			var to2D [][]string

			err := json.Unmarshal(*value, &to1D)
			if err != nil {
				err = nil
				err = json.Unmarshal(*value, &to2D)
				if err != nil {
					continue
				}

				for _, to := range to2D {
					to1D = append(to1D, to...)
				}
			}

			for _, to := range to1D {

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

				a.To = append(a.To, &feed)
			}

			// if lowerKey != "id" && lowerKey != "actor" && lowerKey != "verb" && lowerKey != "object" && lowerKey != "target" && lowerKey != "time" && lowerKey != "foreign_id" && lowerKey != "data" && lowerKey != "to"
		} else {
			var strValue string
			json.Unmarshal(*value, &strValue)
			metadata[key] = strValue
		}
	}

	a.MetaData = metadata
	return nil

}

type postNotificationFeedOutputActivities struct {
	Activities []*NotificationFeedActivity `json:"activities"`
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
			outputResult.Activities = append(outputResult.Activities, activity)
		}

		results = append(results, &outputResult)
	}

	output.Results = results

	return &output
}

type getNotificationFeedOutputResult struct {
	Activities    []*NotificationFeedActivity `json:"activities"`
	ActivityCount int                         `json:"activity_count"`
	ActorCount    int                         `json:"actor_count"`
	CreatedAt     string                      `json:"created_at"`
	Group         string                      `json:"group"`
	ID            string                      `json:"id"`
	IsRead        bool                        `json:"is_read"`
	IsSeen        bool                        `json:"is_seen"`
	UpdatedAt     string                      `json:"updated_at"`
	Verb          string                      `json:"verb"`
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
