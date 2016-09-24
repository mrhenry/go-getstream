package getstream

// ScopeAction defines the Actions allowed by a scope token
type ScopeAction uint32

const (
	// ScopeActionRead : GET, OPTIONS, HEAD
	ScopeActionRead ScopeAction = 1
	// ScopeActionWrite : POST, PUT, PATCH
	ScopeActionWrite ScopeAction = 2
	// ScopeActionDelete : DELETE
	ScopeActionDelete ScopeAction = 4
	// ScopeActionAll : The JWT has permission to all HTTP verbs
	ScopeActionAll ScopeAction = 8
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
	// ScopeContextActivities :  Activities Endpoint
	ScopeContextActivities ScopeContext = 1
	// ScopeContextFeed : Feed Endpoint
	ScopeContextFeed ScopeContext = 2
	// ScopeContextFollower : Following + Followers Endpoint
	ScopeContextFollower ScopeContext = 4
	// ScopeContextAll : Allow access to any resource
	ScopeContextAll ScopeContext = 8
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

// authenticationKind :
type authenticationKind uint32

const (
	// feedAuthentication :
	feedAuthentication authenticationKind = 1
	// appAuthentication :
	appAuthentication authenticationKind = 2
)

// authenticationKind determines if a token or signature is used
type authenticationMethod uint32

const (
	// signatureAuthentication :
	signatureAuthentication authenticationMethod = 1
	// jwtAuthentication :
	jwtAuthentication authenticationMethod = 2
)
