package getstream

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// get request helper
func (c *Client) get(f Feed, path string, signature string, payload []byte, params map[string]string) ([]byte, error) {
	res, err := c.request(f, "GET", path, signature, payload, params)
	return res, err
}

// post request helper
func (c *Client) post(f Feed, path string, signature string, payload []byte, params map[string]string) ([]byte, error) {
	res, err := c.request(f, "POST", path, signature, payload, params)
	return res, err
}

// delete request helper
func (c *Client) del(f Feed, path string, signature string, payload []byte, params map[string]string) error {
	_, err := c.request(f, "DELETE", path, signature, payload, params)
	return err
}

// request helper
func (c *Client) request(f Feed, method string, path string, signature string, payload []byte, params map[string]string) ([]byte, error) {

	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	url = c.baseURL.ResolveReference(url)

	query := url.Query()
	query.Set("api_key", c.Key)
	if c.Location == "" {
		query.Set("location", "unspecified")
	} else {
		query.Set("location", c.Location)
	}

	for key, value := range params {
		query.Set(key, value)
	}
	url.RawQuery = query.Encode()

	// create a new http request
	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// set the Auth headers for the http request
	req.Header.Set("Content-Type", "application/json")
	if f.Token() != "" {
		req.Header.Set("Authorization", signature)
	}

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
		if body != nil {
			return body, nil
		}
		return nil, nil
	default:
		var respErr Error
		err = json.Unmarshal(body, &respErr)
		if err != nil {
			return nil, err
		}
		return nil, &respErr
	}
}
