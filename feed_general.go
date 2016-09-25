package getstream

// GeneralFeed is a container for Feeds returned from request
// The specific Type will be unknown so no Actions are associated with a GeneralFeed
type GeneralFeed struct {
	Client   *Client
	feedSlug string
	userID   string
	token    string
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *GeneralFeed) Signature() string {
	if f.Token() == "" {
		return f.FeedIDWithoutColon()
	}
	return f.FeedIDWithoutColon() + " " + f.Token()
}

// FeedSlug returns the feed slug, it is needed to conform to the Feed interface
func (f *GeneralFeed) FeedSlug() string {
	return f.feedSlug
}

// UserID returns the user id, it is needed to conform to the Feed interface
func (f *GeneralFeed) UserID() string {
	return f.userID
}

// FeedIDWithColon is the combo of the FeedSlug and UserID : "FeedSlug:UserID"
func (f *GeneralFeed) FeedIDWithColon() string {
	return f.FeedSlug() + ":" + f.UserID()
}

// FeedIDWithoutColon is the combo of the FeedSlug and UserID : "FeedSlugUserID"
func (f *GeneralFeed) FeedIDWithoutColon() string {
	return f.FeedSlug() + f.UserID()
}

// SignFeed sets the token on a Feed
func (f *GeneralFeed) SignFeed(signer *Signer) {
	if f.Client.Signer != nil {
		f.token = signer.GenerateToken(f.FeedIDWithoutColon())
	}
}

// Token returns the token of a Feed
func (f *GeneralFeed) Token() string {
	return f.token
}

// GenerateToken returns a new Token for a Feed without setting it to the Feed
func (f *GeneralFeed) GenerateToken(signer *Signer) string {
	if f.Client.Signer != nil {
		return signer.GenerateToken(f.FeedSlug() + f.UserID())
	}
	return ""
}

// Unfollow is used to Unfollow a target Feed
func (f *GeneralFeed) Unfollow(client *Client, target Feed) error {
	f.Client = client
	f.SignFeed(f.Client.Signer)

	endpoint := "feed/" + f.FeedSlug() + "/" + f.UserID() + "/" + "following" + "/" + target.FeedIDWithColon() + "/"

	return f.Client.del(f, endpoint, nil, nil)
}
