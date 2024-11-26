package farcaster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// LocalAccount represents a local wallet account
type LocalAccount struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

// Warpcast represents a client for interacting with the Farcaster API
type Warpcast struct {
	config           *ConfigurationParams
	wallet           *LocalAccount
	accessToken      *string
	expiresAt        *int64
	rotationDuration int64
	client           *http.Client
	baseHeaders      map[string]string
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

// NewWarpcast creates a new Warpcast client instance
func NewWarpcast(opts ...WarpcastOption) (*Warpcast, error) {
	w := &Warpcast{
		config: &ConfigurationParams{
			BasePath: "https://api.warpcast.com/v2/",
		},
		rotationDuration: 10,
		client:           &http.Client{},
		baseHeaders:      make(map[string]string),
	}

	// Apply options
	for _, opt := range opts {
		opt(w)
	}

	// Validate and setup authentication
	if w.accessToken != nil {
		w.baseHeaders["Authorization"] = fmt.Sprintf("Bearer %s", *w.accessToken)
		if w.expiresAt == nil {
			// Set to year 3000
			future := int64(33228645430000)
			w.expiresAt = &future
		}
	} else if w.wallet == nil {
		return nil, fmt.Errorf("no wallet or access token provided")
	} else {
		if err := w.createNewAuthToken(w.rotationDuration); err != nil {
			return nil, fmt.Errorf("failed to create auth token: %w", err)
		}
	}

	return w, nil
}

// WarpcastOption defines a function type for configuring the Warpcast client
type WarpcastOption func(*Warpcast)

// WithAccessToken sets the access token for the client
func WithAccessToken(token string, expiresAt *int64) WarpcastOption {
	return func(w *Warpcast) {
		w.accessToken = &token
		w.expiresAt = expiresAt
	}
}

// WithWallet sets the wallet for the client
func WithWallet(wallet *LocalAccount) WarpcastOption {
	return func(w *Warpcast) {
		w.wallet = wallet
	}
}

// request performs an HTTP request and handles the response
func (w *Warpcast) request(method, path string, params map[string]string, body interface{}, headers map[string]string) ([]byte, error) {
	if err := w.checkAuthHeader(); err != nil {
		return nil, fmt.Errorf("auth check failed: %w", err)
	}

	url := w.config.BasePath + path

	// Create request
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add query parameters
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	// Add headers
	for k, v := range w.baseHeaders {
		req.Header.Set(k, v)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	// Make request
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for API errors
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err == nil {
		if errors, ok := result["errors"]; ok {
			return nil, fmt.Errorf("API error: %v", errors)
		}
	}

	return respBody, nil
}

// GetHealthcheck checks if the API is up and running
func (w *Warpcast) GetHealthcheck() (bool, error) {
	resp, err := w.client.Get("https://api.warpcast.com/healthcheck")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK, nil
}

// checkAuthHeader verifies and refreshes the authentication token if needed
func (w *Warpcast) checkAuthHeader() error {
	if w.expiresAt == nil {
		return fmt.Errorf("expires_at is not set")
	}

	if *w.expiresAt < nowMs()+1000 {
		if err := w.createNewAuthToken(w.rotationDuration); err != nil {
			return fmt.Errorf("failed to refresh auth token: %w", err)
		}
	}
	return nil
}

// nowMs returns the current time in milliseconds
func nowMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// createNewAuthToken generates a new authentication token
func (w *Warpcast) createNewAuthToken(duration int64) error {
	// TODO: Implement actual token creation logic
	// This is a placeholder implementation
	token := "temporary_token"
	expiresAt := nowMs() + (duration * 60 * 1000) // Convert minutes to milliseconds

	w.accessToken = &token
	w.expiresAt = &expiresAt
	w.baseHeaders["Authorization"] = fmt.Sprintf("Bearer %s", token)

	return nil
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

// GetAsset retrieves asset information
func (w *Warpcast) GetAsset(tokenID int) (*AssetResult, error) {
	params := map[string]string{
		"token_id": fmt.Sprintf("%d", tokenID),
	}

	resp, err := w.request("GET", "asset", params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}

	var result struct {
		Result AssetResult `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal asset response: %w", err)
	}

	return &result.Result, nil
}

// GetAssetEvents retrieves events for a given asset
func (w *Warpcast) GetAssetEvents(cursor *string, limit int) (*IterableEventsResult, error) {
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	if cursor != nil {
		params["cursor"] = *cursor
	}

	resp, err := w.request("GET", "asset-events", params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset events: %w", err)
	}

	var result struct {
		Result struct {
			Events []Event `json:"events"`
		} `json:"result"`
		Next *struct {
			Cursor string `json:"cursor"`
		} `json:"next"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal asset events response: %w", err)
	}

	var nextCursor *string
	if result.Next != nil {
		nextCursor = &result.Next.Cursor
	}

	return &IterableEventsResult{
		Events: result.Result.Events,
		Cursor: nextCursor,
	}, nil
}

// PutAuth generates a custody bearer token and uses it to generate an access token
func (w *Warpcast) PutAuth(authParams *AuthParams) (*TokenResult, error) {
	header, err := w.generateCustodyAuthHeader(authParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate auth header: %w", err)
	}

	body := struct {
		Params *AuthParams `json:"params"`
	}{
		Params: authParams,
	}

	headers := map[string]string{
		"Authorization": header,
	}

	resp, err := w.request("PUT", "auth", nil, body, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to put auth: %w", err)
	}

	var result struct {
		Result TokenResult `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal auth response: %w", err)
	}

	return &result.Result, nil
}

// DeleteAuth deletes an access token
func (w *Warpcast) DeleteAuth() (*StatusContent, error) {
	timestamp := nowMs()
	body := struct {
		Params struct {
			Timestamp int64 `json:"timestamp"`
		} `json:"params"`
	}{
		Params: struct {
			Timestamp int64 `json:"timestamp"`
		}{
			Timestamp: timestamp,
		},
	}

	resp, err := w.request("DELETE", "auth", nil, body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete auth: %w", err)
	}

	var result struct {
		Result StatusContent `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal delete auth response: %w", err)
	}

	return &result.Result, nil
}

// generateCustodyAuthHeader generates the custody authorization header
func (w *Warpcast) generateCustodyAuthHeader(authParams *AuthParams) (string, error) {
	if w.wallet == nil {
		return "", fmt.Errorf("wallet is required for custody auth")
	}
	// TODO: Implement actual custody signature logic
	// This is a placeholder that should be replaced with proper signature generation
	return fmt.Sprintf("Bearer %s", w.wallet.PrivateKey), nil
}

// ReactionsPutResult represents the result of liking a cast
type ReactionsPutResult struct {
	Like struct {
		CastHash   string `json:"castHash"`
		ReactorFid int    `json:"reactorFid"`
		Timestamp  int64  `json:"timestamp"`
	} `json:"like"`
}

// LikeCast likes a given cast
func (w *Warpcast) LikeCast(castHash string) (*ReactionsPutResult, error) {
	body := struct {
		CastHash string `json:"castHash"`
	}{
		CastHash: castHash,
	}

	resp, err := w.request("PUT", "cast-likes", nil, body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to like cast: %w", err)
	}

	var result struct {
		Result ReactionsPutResult `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal like cast response: %w", err)
	}

	return &result.Result, nil
}
