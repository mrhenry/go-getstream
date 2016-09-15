package getstream

// GeneralFeed is a container for Feeds returned from request
// The specific Type will be unknown so no Actions are associated with a GeneralFeed
type GeneralFeed struct {
	client   *Client
	FeedSlug string
	UserID   string
	token    string
}

// Client returns the Client associated with the GeneralFeed
func (f GeneralFeed) Client() *Client {
	return f.client
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *GeneralFeed) Signature() string {
	if f.Token() == "" {
		return f.feedIDWithoutColon()
	}
	return f.feedIDWithoutColon() + " " + f.Token()
}

// FeedID is the combo if the FeedSlug and UserID : "FeedSlug:UserID"
func (f *GeneralFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

func (f *GeneralFeed) feedIDWithoutColon() string {
	return f.FeedSlug + f.UserID
}

// Token returns the token of a Feed
func (f *GeneralFeed) Token() string {
	return f.token
}

// GenerateToken returns a new Token for a Feed without setting it to the Feed
func (f *GeneralFeed) GenerateToken(signer *Signer) string {
	if f.Client().Signer != nil {
		return signer.generateToken(f.FeedSlug + f.UserID)
	}
	return ""
}

// Unfollow is used to Unfollow a target Feed
func (f *GeneralFeed) Unfollow(client *Client, target *FlatFeed) error {

	f.client = client
	f.token = f.GenerateToken(f.Client().Signer)

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + target.FeedID().Value() + "/"

	return f.Client().del(f, endpoint, nil, nil)

}
