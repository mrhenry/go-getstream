package getstream

// ScopeAction defines the Actions allowed by a scope token
type ScopeAction uint32

const (
	// ReadAction : GET, OPTIONS, HEAD
	ReadAction ScopeAction = 1
	// WriteAction : POST, PUT, PATCH
	WriteAction ScopeAction = 2
	// DeleteAction : DELETE
	DeleteAction ScopeAction = 4
	// AllActions : The JWT has permission to all HTTP verbs
	AllActions ScopeAction = 8
)

// Value returns a string representation
func (a ScopeAction) Value() string {
	switch a {
	case 1:
		return "read"
	case 2:
		return "write"
	case 4:
		return "delete"
	case 8:
		return "*"
	default:
		return ""
	}
}

// ScopeContext defines the resources accessible by a scope token
type ScopeContext uint32

const (
	// ActivitiesContext :  Activities Endpoint
	ActivitiesContext ScopeContext = 1
	// FeedContext : Feed Endpoint
	FeedContext ScopeContext = 2
	// FollowerContext : Following + Followers Endpoint
	FollowerContext ScopeContext = 4
	// AllContexts : Allow access to any resource
	AllContexts ScopeContext = 8
)

// Value returns a string representation
func (a ScopeContext) Value() string {
	switch a {
	case 1:
		return "activities"
	case 2:
		return "feed"
	case 4:
		return "follower"
	case 8:
		return "*"
	default:
		return ""
	}
}
