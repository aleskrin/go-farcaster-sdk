package farcaster

import (
	"encoding/json"
	"fmt"
)

// CastsResult represents a collection of casts
type CastsResult struct {
	Casts []CastContent `json:"casts"`
}

// GetAllCastsInThread retrieves all casts in a thread
// Parameters:
//   - threadHash: The hash of the thread to retrieve casts for
//
// Returns:
//   - *CastsResult: A collection of casts
//   - error: Any error that occurred
func (w *Warpcast) GetAllCastsInThread(threadHash string) (*CastsResult, error) {
	params := map[string]string{
		"threadHash": threadHash,
	}

	resp, err := w.request("GET", "all-casts-in-thread", params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get thread casts: %w", err)
	}

	var result struct {
		Result CastsResult `json:"result"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal thread casts response: %w", err)
	}

	return &result.Result, nil
}
