package getstream

// Option is a mutator for Client.
// Use it to construct a Client with custom values.
type Option func(*Client)

func ServerOptions(key string, secret string, appID string, location string) Option {
	return func(c *Client) {
		c.Key = key
		c.Secret = secret
		c.AppID = appID
		c.Location = location
	}
}
