package farcaster

import (
	"encoding/json"
	"fmt"
)

// FollowUser follows a user with the given FID
//
// Parameters:
//   - fid: Farcaster ID of the user to follow
//
// Returns:
//   - *StatusContent: Status of the follow operation
//   - error: Any error that occurred
func (w *Warpcast) FollowUser(fid int) (*StatusContent, error) {
	body := FollowsPutRequest{
		TargetFid: fid,
	}

	resp, err := w.request("PUT", "follows", nil, body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to follow user: %w", err)
	}

	var result StatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal follow response: %w", err)
	}

	return &result.Result, nil
} 