package getstream

type postNotificationFeedOutputActivities struct {
	Activities []*Activity `json:"activities"`
}

// GetNotificationFeedInput is used to Get a list of Activities from a NotificationFeed
type GetNotificationFeedInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE string `json:"id_gte,omitempty"`
	IDGT  string `json:"id_gt,omitempty"`
	IDLTE string `json:"id_lte,omitempty"`
	IDLT  string `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}

// GetNotificationFeedOutput is the response from a NotificationFeed Activities Get Request
type GetNotificationFeedOutput struct {
	Duration string
	Next     string
	Results  []*struct {
		Activities    []*Activity
		ActivityCount int
		ActorCount    int
		CreatedAt     string
		Group         string
		ID            string
		IsRead        bool
		IsSeen        bool
		UpdatedAt     string
		Verb          string
	}
	Unread int
	Unseen int
}

type getNotificationFeedOutput struct {
	Duration string                             `json:"duration"`
	Next     string                             `json:"next"`
	Results  []*getNotificationFeedOutputResult `json:"results"`
	Unread   int                                `json:"unread"`
	Unseen   int                                `json:"unseen"`
}

func (a getNotificationFeedOutput) output() *GetNotificationFeedOutput {

	output := GetNotificationFeedOutput{
		Duration: a.Duration,
		Next:     a.Next,
		Unread:   a.Unread,
		Unseen:   a.Unseen,
	}

	var results []*struct {
		Activities    []*Activity
		ActivityCount int
		ActorCount    int
		CreatedAt     string
		Group         string
		ID            string
		IsRead        bool
		IsSeen        bool
		UpdatedAt     string
		Verb          string
	}

	for _, result := range a.Results {

		outputResult := struct {
			Activities    []*Activity
			ActivityCount int
			ActorCount    int
			CreatedAt     string
			Group         string
			ID            string
			IsRead        bool
			IsSeen        bool
			UpdatedAt     string
			Verb          string
		}{
			ActivityCount: result.ActivityCount,
			ActorCount:    result.ActorCount,
			CreatedAt:     result.CreatedAt,
			Group:         result.Group,
			ID:            result.ID,
			IsRead:        result.IsRead,
			IsSeen:        result.IsSeen,
			UpdatedAt:     result.UpdatedAt,
			Verb:          result.Verb,
		}

		for _, activity := range result.Activities {
			outputResult.Activities = append(outputResult.Activities, activity)
		}

		results = append(results, &outputResult)
	}

	output.Results = results

	return &output
}

type getNotificationFeedOutputResult struct {
	Activities    []*Activity `json:"activities"`
	ActivityCount int         `json:"activity_count"`
	ActorCount    int         `json:"actor_count"`
	CreatedAt     string      `json:"created_at"`
	Group         string      `json:"group"`
	ID            string      `json:"id"`
	IsRead        bool        `json:"is_read"`
	IsSeen        bool        `json:"is_seen"`
	UpdatedAt     string      `json:"updated_at"`
	Verb          string      `json:"verb"`
}

type getNotificationFeedFollowersInput struct {
	Limit int `json:"limit"`
	Skip  int `json:"offset"`
}

type getNotificationFeedFollowersOutput struct {
	Duration string                                      `json:"duration"`
	Results  []*getNotificationFeedFollowersOutputResult `json:"results"`
}

type getNotificationFeedFollowersOutputResult struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	FeedID    string `json:"feed_id"`
	TargetID  string `json:"target_id"`
}

type postNotificationFeedFollowingInput struct {
	Target            string `json:"target"`
	ActivityCopyLimit int    `json:"activity_copy_limit"`
}
