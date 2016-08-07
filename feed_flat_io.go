package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"
)

// FlatFeedActivity is a getstream Activity
// Use it to post activities to FlatFeeds
// It is also the response from FlatFeed Fetch and List Requests
type FlatFeedActivity struct {
	ID        string
	Actor     FeedID
	Verb      string
	Object    FeedID
	Target    FeedID
	TimeStamp *time.Time

	ForeignID string
	Data      json.RawMessage
	MetaData  map[string]string

	To []Feed
}

func (a FlatFeedActivity) MarshalJSON() ([]byte, error) {

	payload := make(map[string]interface{})

	for key, value := range a.MetaData {
		payload[key] = value
	}

	payload["actor"] = string(a.Actor)
	payload["verb"] = string(a.Verb)
	payload["object"] = string(a.Object)

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

func (a *FlatFeedActivity) UnmarshalJSON(b []byte) (err error) {

	rawPayload := make(map[string]json.RawMessage)
	metadata := make(map[string]string)

	json.Unmarshal(b, &rawPayload)

	for key, value := range rawPayload {
		lowerKey := strings.ToLower(key)

		if lowerKey == "id" {
			var strValue string
			json.Unmarshal(value, &strValue)
			a.ID = strValue
		} else if lowerKey == "actor" {
			var strValue string
			json.Unmarshal(value, &strValue)
			a.Actor = FeedID(strValue)
		} else if lowerKey == "verb" {
			var strValue string
			json.Unmarshal(value, &strValue)
			a.Verb = strValue
		} else if lowerKey == "foreign_id" {
			var strValue string
			json.Unmarshal(value, &strValue)
			a.ForeignID = strValue
		} else if lowerKey == "object" {
			var strValue string
			json.Unmarshal(value, &strValue)
			a.Object = FeedID(strValue)
		} else if lowerKey == "target" {
			var strValue string
			json.Unmarshal(value, &strValue)
			a.Target = FeedID(strValue)
		} else if lowerKey == "time" {
			var strValue string
			err := json.Unmarshal(value, &strValue)
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

			err := json.Unmarshal(value, &to1D)
			if err != nil {
				err = nil
				err = json.Unmarshal(value, &to2D)
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
			json.Unmarshal(value, &strValue)
			metadata[key] = strValue
		}
	}

	a.MetaData = metadata
	return nil

}

type postFlatFeedOutputActivities struct {
	Activities []*FlatFeedActivity `json:"activities"`
}

// GetFlatFeedInput is used to Get a list of Activities from a FlatFeed
type GetFlatFeedInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE int `json:"id_gte,omitempty"`
	IDGT  int `json:"id_gt,omitempty"`
	IDLTE int `json:"id_lte,omitempty"`
	IDLT  int `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}

type GetFlatFeedOutput struct {
	Duration   string              `json:"duration"`
	Next       string              `json:"next"`
	Activities []*FlatFeedActivity `json:"results"`
}

type getFlatFeedFollowersInput struct {
	Limit int `json:"limit"`
	Skip  int `json:"offset"`
}

type getFlatFeedFollowersOutput struct {
	Duration string                              `json:"duration"`
	Results  []*getFlatFeedFollowersOutputResult `json:"results"`
}

type getFlatFeedFollowersOutputResult struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	FeedID    string `json:"feed_id"`
	TargetID  string `json:"target_id"`
}

type postFlatFeedFollowingInput struct {
	Target            string `json:"target"`
	ActivityCopyLimit int    `json:"activity_copy_limit"`
}
