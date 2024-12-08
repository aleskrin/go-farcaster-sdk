package farcaster

import (
	"encoding/json"
	"fmt"
)

// UnfollowUser unfollows a user with the given FID
//
// Parameters:
//   - fid: Farcaster ID of the user to unfollow
//
// Returns:
//   - *StatusContent: Status of the unfollow operation
//   - error: Any error that occurred
func (w *Warpcast) UnfollowUser(fid int) (*StatusContent, error) {
	body := FollowsDeleteRequest{
		TargetFid: fid,
	}

	resp, err := w.request("DELETE", "follows", nil, body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to unfollow user: %w", err)
	}

	var result StatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unfollow response: %w", err)
	}

	return &result.Result, nil
} 