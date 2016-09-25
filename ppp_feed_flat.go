package getstream

type postFlatFeedOutputActivities struct {
	Activities []*Activity `json:"activities"`
}

// GetFlatFeedOutput is the response from a FlatFeed Activities Get Request
type GetFlatFeedOutput struct {
	Duration   string      `json:"duration"`
	Next       string      `json:"next"`
	Activities []*Activity `json:"results"`
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
