package farcaster

import "net/http"

// Warpcast represents a client for interacting with the Farcaster API
type Warpcast struct {
	config           *ConfigurationParams
	wallet           *LocalAccount
	accessToken      *string
	expiresAt        *int64
	rotationDuration int64
	client           *http.Client
	baseHeaders      map[string]string
	BaseURL          string
	HTTPClient       *http.Client
}

// LocalAccount represents a local wallet account
type LocalAccount struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

// ConfigurationParams holds the configuration for the Warpcast client
type ConfigurationParams struct {
	BasePath    string
	BaseOptions map[string]interface{}
	Username    *string
}

// AuthParams represents the parameters needed for authentication
type AuthParams struct {
	Timestamp int64 `json:"timestamp"`
	// Add other required fields based on the API documentation
}

// AssetResult represents the result of an asset query
type AssetResult struct {
	// Add fields as needed
}

// TokenResult represents the result of a token operation
type TokenResult struct {
	Token struct {
		Secret string `json:"secret"`
	} `json:"token"`
}

// StatusContent represents a status response
type StatusContent struct {
	Success bool `json:"success"`
}

// Event represents a single asset event
type Event struct {
	// Add relevant fields based on your API response
}

// IterableEventsResult represents a paginated list of events
type IterableEventsResult struct {
	Events []Event
	Cursor *string
}

// ReactionsPutResult represents the result of liking a cast
type ReactionsPutResult struct {
	Like struct {
		CastHash   string `json:"castHash"`
		ReactorFid int    `json:"reactorFid"`
		Timestamp  int64  `json:"timestamp"`
	} `json:"like"`
}

// CastContent represents the content of a cast
type CastContent struct {
	Hash      string    `json:"hash"`
	ThreadHash string   `json:"threadHash"`
	ParentHash *string  `json:"parentHash,omitempty"`
	Author    Author    `json:"author"`
	Text      string    `json:"text"`
	Embeds    []string  `json:"embeds,omitempty"`
	// Add other fields as needed based on the API response
}

type Author struct {
	FID       int    `json:"fid"`
	Username  string `json:"username"`
	DisplayName string `json:"displayName"`
	// Add other author fields as needed
}

// Parent represents the parent of a cast
type Parent struct {
	Hash string `json:"hash"`
}

// CastsPostRequest represents the request body for posting a cast
type CastsPostRequest struct {
	Text       string   `json:"text"`
	Embeds     []string `json:"embeds,omitempty"`
	Parent     *Parent  `json:"parent,omitempty"`
	ChannelKey *string  `json:"channelKey,omitempty"`
}

// CastsPostResponse represents the response from posting a cast
type CastsPostResponse struct {
	Result CastContent `json:"result"`
}

// Add these new types to types.go

type UsersResult struct {
	Users []ApiUser `json:"users"`
}

type FollowingGetResponse struct {
	Result struct {
		Users []ApiUser `json:"users"`
	} `json:"result"`
	Next *struct {
		Cursor string `json:"cursor"`
	} `json:"next"`
}

// FollowsPutRequest represents the request to follow a user
type FollowsPutRequest struct {
	TargetFid int `json:"targetFid"`
}

// StatusResponse represents the response from a follow operation
type StatusResponse struct {
	Result StatusContent `json:"result"`
}
