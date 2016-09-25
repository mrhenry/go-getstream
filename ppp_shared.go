package getstream

// GetActivitiesInput is used to Get a list of Activities from a Feed
type GetActivitiesInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE string `json:"id_gte,omitempty"`
	IDGT  string `json:"id_gt,omitempty"`
	IDLTE string `json:"id_lte,omitempty"`
	IDLT  string `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}
