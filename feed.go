package getstream

// Feed is the interface bundling all Feed Types
// It exposes methods needed for all Types
type Feed interface {
	Signature() string
	FeedSlug() string
	UserID() string
	FeedIDWithoutColon() string
	FeedIDWithColon() string
	Token() string
	SignFeed(signer *Signer)
	GenerateToken(signer *Signer) string
}

type postFeedFollowingManyInput struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
