package getstream

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const VERSION="v1.0.0"

// Client is used to connect to getstream.io
type Client struct {
	HTTP    *http.Client
	BaseURL *url.URL // https://api.getstream.io/api/
	Config  *Config
	Signer  *Signer
}

/**
 * New returns a GetStream client.
 *
 * Params:
 *   cfg, pointer to a Config structure which takes the API credentials, Location, etc
 * Returns:
 *   Client struct
 */
func New(cfg *Config) (*Client, error) {
	var (
		timeout int64
	)

	if cfg.APIKey == "" {
		return nil, errors.New("Required API Key was not set")
	}

	if cfg.APISecret == "" && cfg.Token == "" {
		return nil, errors.New("API Secret or Token was not set, one or the other is required")
	}

	if cfg.TimeoutInt <= 0 {
		timeout = 3
	} else {
		timeout = cfg.TimeoutInt
	}
	cfg.SetTimeout(timeout)

	if cfg.Version == "" {
		cfg.Version = "v1.0"
	}

	location := "api"
	if cfg.Location != "" {
		location = cfg.Location + "-api"
	}

	baseURL, err := url.Parse("https://" + location + ".getstream.io/api/" + cfg.Version + "/")
	if err != nil {
		return nil, err
	}
	cfg.SetBaseURL(baseURL)

	var signer *Signer
	if cfg.Token != "" {
		// build the Signature mechanism based on a Token value passed to the client setup
		cfg.SetAPISecret("")
		signer = &Signer{
			Secret: cfg.Token,
		}
	} else {
		// build the Signature based on the API Secret
		cfg.SetToken("")
		signer = &Signer{
			Secret: cfg.APISecret,
		}
	}

	client := &Client{
		HTTP: &http.Client{
			Timeout: cfg.TimeoutDuration,
		},
		BaseURL: baseURL,
		Config:  cfg,
		Signer:  signer,
	}

	return client, nil
}

// FlatFeed returns a getstream feed
// Slug is the FlatFeed Group name
// id is the Specific FlatFeed inside a FlatFeed Group
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) FlatFeed(feedSlug string, userID string) (*FlatFeed, error) {
	var err error

	feedSlug, err = ValidateFeedSlug(feedSlug)
	if err != nil {
		return nil, err
	}
	userID, err = ValidateUserID(userID)
	if err != nil {
		return nil, err
	}

	feed := &FlatFeed{
		Client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	feed.SignFeed(c.Signer)
	return feed, nil
}

// NotificationFeed returns a getstream feed
// Slug is the NotificationFeed Group name
// id is the Specific NotificationFeed inside a NotificationFeed Group
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) NotificationFeed(feedSlug string, userID string) (*NotificationFeed, error) {
	var err error

	feedSlug, err = ValidateFeedSlug(feedSlug)
	if err != nil {
		return nil, err
	}
	userID, err = ValidateUserID(userID)
	if err != nil {
		return nil, err
	}

	feed := &NotificationFeed{
		Client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	feed.SignFeed(c.Signer)
	return feed, nil
}

// AggregatedFeed returns a getstream feed
// Slug is the AggregatedFeed Group name
// id is the Specific AggregatedFeed inside a AggregatedFeed Group
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) AggregatedFeed(feedSlug string, userID string) (*AggregatedFeed, error) {
	var err error

	feedSlug, err = ValidateFeedSlug(feedSlug)
	if err != nil {
		return nil, err
	}
	userID, err = ValidateUserID(userID)
	if err != nil {
		return nil, err
	}

	feed := &AggregatedFeed{
		Client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	feed.SignFeed(c.Signer)
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
func (c *Client) AbsoluteURL(path string) (*url.URL, error) {

	result, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// DEBUG: Use this line to send stuff to a proxy instead.
	// c.baseURL, _ = url.Parse("http://0.0.0.0:8000/")
	result = c.BaseURL.ResolveReference(result)

	qs := result.Query()
	qs.Set("api_key", c.Config.APIKey)
	if c.Config.Location == "" {
		qs.Set("location", "unspecified")
	} else {
		qs.Set("location", c.Config.Location)
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
