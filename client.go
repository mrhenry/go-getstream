package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"time"
	"gopkg.in/LeisureLink/httpsig.v1"
)


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
	port := ""
	secure := "s"
	if cfg.Location != "" {
		location = cfg.Location + "-api"
		if cfg.Location == "localhost" {
			port = ":8000"
			secure = ""
		}
	}

	baseURL, err := url.Parse("http" + secure + "://" + location + ".getstream.io" + port + "/api/" + cfg.Version + "/")
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
	if c.Config.Location == "" || c.Config.Location == "localhost" {
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

// get request helper
func (c *Client) get(f Feed, path string, payload []byte, params map[string]string) ([]byte, error) {
	res, err := c.request(f, "GET", path, payload, params)
	return res, err
}

// post request helper
func (c *Client) post(f Feed, path string, payload []byte, params map[string]string) ([]byte, error) {
	fmt.Println(f, "POST", path, string(payload), params)
	res, err := c.request(f, "POST", path, payload, params)
	return res, err
}

// delete request helper
func (c *Client) del(f Feed, path string, payload []byte, params map[string]string) error {
	_, err := c.request(f, "DELETE", path, payload, params)
	return err
}

// request helper
func (c *Client) request(f Feed, method string, path string, payload []byte, params map[string]string) ([]byte, error) {
	apiUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	apiUrl = c.BaseURL.ResolveReference(apiUrl)

	query := apiUrl.Query()
	query = c.setStandardParams(query)
	query = c.setRequestParams(query, params)
	apiUrl.RawQuery = query.Encode()

	// create a new http request
	req, err := http.NewRequest(method, apiUrl.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// set the Auth headers for the http request
	c.setHeaders(req, f)

	fmt.Println("req.Body", req.Body)
	fmt.Println("req.Header", req.Header)
	fmt.Println("req.URL", req.URL)

	// perform the http request
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle the response
	switch {
	case resp.StatusCode/100 == 2: // SUCCESS
		fmt.Println("200 success:", string(body))
		return body, nil
	default:
		var respErr Error
		err = json.Unmarshal(body, &respErr)
		if err != nil {
			fmt.Println("client.request non-200 success, error unmarshalling json response, body:", body)
			return nil, err
		}
		return nil, &respErr
	}
}

func (c *Client) setStandardParams(query url.Values) url.Values {
	query.Set("api_key", c.Config.APIKey)
	if c.Config.Location == "" || c.Config.Location == "localhost" {
		query.Set("location", "unspecified")
	} else {
		query.Set("location", c.Config.Location)
	}

	return query
}

func (c *Client) setRequestParams(query url.Values, params map[string]string) url.Values {
	for key, value := range params {
		query.Set(key, value)
	}
	return query
}

func (c *Client) setHeaders(request *http.Request, f Feed) error {
	request.Header.Set("X-Stream-Client", "stream-go-client-" + VERSION)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Api-Key", c.Config.APIKey)

	t := time.Now()
	request.Header.Set("Date", t.Format("Mon, 2 Jan 2006 15:04:05 MST"))

	fmt.Println("is feed nil?", f)
	fmt.Println("c.Config.APISecret:", c.Config.APISecret)
	fmt.Println("c.Config.Token:", c.Config.Token)
	if f != nil {
		fmt.Println("f.Token:", f.Token())
		fmt.Println("f.Signature:", f.Signature())
	}

	if f != nil && c.Config.APISecret != "" && f.Token() != "" {
		fmt.Println(1)
		request.Header.Set("Authorization", f.Signature())
		return nil
	} else if c.Config.APISecret != "" && (f == nil || f.Token() == "") {
		fmt.Println(2)
		signer, _ := httpsig.NewRequestSigner(c.Config.APIKey, c.Config.APISecret, "hmac-sha256")
		signer.SignRequest(request, []string{}, nil)
		return nil
	} else if c.Config.Token != "" {
		fmt.Println(3)
		request.Header.Set("stream-auth-type", "jwt")
		request.Header.Set("Authorization", c.Config.Token)
		return nil
	}

	return errors.New("No API Secret or config/feed Token")

}

type PostFlatFeedFollowingManyInput struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
}
/** PrepFollowFlatFeed - prepares JSON needed for one feed to follow another

	Params:
	targetFeed, FlatFeed which wants to follow another
	sourceFeed, FlatFeed which is to be followed

	Returns:
	[]byte, array of bytes of JSON suitable for API consumption
 */
func (c *Client) PrepFollowFlatFeed(targetFeed *FlatFeed, sourceFeed *FlatFeed) *PostFlatFeedFollowingManyInput {
	return &PostFlatFeedFollowingManyInput{
		Source: sourceFeed.FeedSlug + ":" + sourceFeed.UserID,
		Target: targetFeed.FeedSlug + ":" + targetFeed.UserID,
	}
}
func (c *Client) PrepFollowAggregatedFeed(targetFeed *FlatFeed, sourceFeed *AggregatedFeed) *PostFlatFeedFollowingManyInput {
	return &PostFlatFeedFollowingManyInput{
		Source: sourceFeed.FeedSlug + ":" + sourceFeed.UserID,
		Target: targetFeed.FeedSlug + ":" + targetFeed.UserID,
	}
}
func (c *Client) PrepFollowNotifcationFeed(targetFeed *FlatFeed, sourceFeed *NotificationFeed) *PostFlatFeedFollowingManyInput {
	return &PostFlatFeedFollowingManyInput{
		Source: sourceFeed.FeedSlug + ":" + sourceFeed.UserID,
		Target: targetFeed.FeedSlug + ":" + targetFeed.UserID,
	}
}

type PostActivityToManyInput struct {
	Activity   Activity `json:"activity"`
	FeedIDs   []string `json:"feeds"`
}

func (c *Client) AddActivityToMany(activity Activity, feeds []string) error {
	endpoint := "feed/add_to_many/"
	params := map[string]string{}

	payload := &PostActivityToManyInput{
		Activity: activity,
		FeedIDs: feeds,
	}

	final_payload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = c.post(nil, endpoint, final_payload, params)
	return err
}