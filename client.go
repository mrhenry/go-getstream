package getstream

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Client is used to connect to getstream.io
type Client struct {
	HTTP    *http.Client
	baseURL *url.URL // https://api.getstream.io/api/

	Key      string
	Secret   string
	AppID    string
	Location string // https://location-api.getstream.io/api/

	Token string

	Signer *Signer
}

// New returns a getstream client.
// Params :
// - api key
// - api secret
// - appID
// - region
// An http.Client with custom settings can be assigned after construction
func New(key string, secret string, appID string, location string) (*Client, error) {
	baseURLStr := "https://api.getstream.io/api/v1.0/"
	if location != "" {
		baseURLStr = "https://" + location + "-api.getstream.io/api/v1.0/"
	}

	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		HTTP: &http.Client{
			Timeout: 3 * time.Second,
		},
		baseURL: baseURL,

		Key:      key,
		Secret:   secret,
		AppID:    appID,
		Location: location,

		Signer: &Signer{
			Secret: secret,
		},
	}, nil
}

// NewWithToken returns a getstream client.
// Params :
// - api key
// - token
// - appID
// - region
// An http.Client with custom settings can be assigned after construction
func NewWithToken(key string, token string, appID string, location string) (*Client, error) {
	baseURLStr := "https://api.getstream.io/api/v1.0/"
	if location != "" {
		baseURLStr = "https://" + location + "-api.getstream.io/api/v1.0/"
	}

	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		HTTP: &http.Client{
			Timeout: 3 * time.Second,
		},
		baseURL: baseURL,

		Key:      key,
		Token:    token,
		AppID:    appID,
		Location: location,
	}, nil
}

// FlatFeed returns a getstream feed
// Slug is the FlatFeed Group name
// id is the Specific FlatFeed inside a FlatFeed Group
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) FlatFeed(feedSlug string, userID string) (*FlatFeed, error) {

	r, err := regexp.Compile(`^\w+$`)
	if err != nil {
		return nil, err
	}

	feedSlug = strings.Replace(feedSlug, "-", "_", -1)
	userID = strings.Replace(userID, "-", "_", -1)

	if !r.MatchString(feedSlug) || !r.MatchString(userID) {
		return nil, errors.New("invalid feedSlug or userID")
	}

	feed := &FlatFeed{
		client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	if c.Signer != nil {
		feed.token = c.Signer.generateToken(feed.feedIDWithoutColon())
	}
	return feed, nil
}

// NotificationFeed returns a getstream feed
// Slug is the NotificationFeed Group name
// id is the Specific NotificationFeed inside a NotificationFeed Group
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) NotificationFeed(feedSlug string, userID string) (*NotificationFeed, error) {

	r, err := regexp.Compile(`^\w+$`)
	if err != nil {
		return nil, err
	}

	feedSlug = strings.Replace(feedSlug, "-", "_", -1)
	userID = strings.Replace(userID, "-", "_", -1)

	if !r.MatchString(feedSlug) || !r.MatchString(userID) {
		return nil, errors.New("invalid feedSlug or userID")
	}

	feed := &NotificationFeed{
		client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	if c.Signer != nil {
		feed.token = c.Signer.generateToken(feed.feedIDWithoutColon())
	}
	return feed, nil
}

// AggregatedFeed returns a getstream feed
// Slug is the AggregatedFeed Group name
// id is the Specific AggregatedFeed inside a AggregatedFeed Group
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) AggregatedFeed(feedSlug string, userID string) (*AggregatedFeed, error) {

	r, err := regexp.Compile(`^\w+$`)
	if err != nil {
		return nil, err
	}

	feedSlug = strings.Replace(feedSlug, "-", "_", -1)
	userID = strings.Replace(userID, "-", "_", -1)

	if !r.MatchString(feedSlug) || !r.MatchString(userID) {
		return nil, errors.New("invalid feedSlug or userID")
	}

	feed := &AggregatedFeed{
		client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	if c.Signer != nil {
		feed.token = c.Signer.generateToken(feed.feedIDWithoutColon())
	}
	return feed, nil
}

// // UpdateActivities is used to update multiple Activities
// func (c *Client) UpdateActivities(activities []interface{}) ([]*Activity, error) {
//
// 	payload, err := json.Marshal(activities)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	endpoint := "activities/"
//
// 	resultBytes, err := c.post(endpoint, payload, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	output := &postFlatFeedOutputActivities{}
// 	err = json.Unmarshal(resultBytes, output)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return output.Activities, err
// }

// absoluteUrl create a url.URL instance and sets query params (bad!!!)
func (c *Client) absoluteURL(path string) (*url.URL, error) {

	result, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// DEBUG: Use this line to send stuff to a proxy instead.
	// c.baseURL, _ = url.Parse("http://0.0.0.0:8000/")
	result = c.baseURL.ResolveReference(result)

	qs := result.Query()
	qs.Set("api_key", c.Key)
	if c.Location == "" {
		qs.Set("location", "unspecified")
	} else {
		qs.Set("location", c.Location)
	}
	result.RawQuery = qs.Encode()

	return result, nil
}

// ConvertUUIDToWord replaces - with _
// It is used by go-getstream to convert UUID to a string that matches the word regex
// You can use it to convert UUID's to match go-getstream internals.
func ConvertUUIDToWord(uuid string) string {

	return strings.Replace(uuid, "-", "_", -1)

}
