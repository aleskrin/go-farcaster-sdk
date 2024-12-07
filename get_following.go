package farcaster

import (
	"encoding/json"
	"fmt"
)

// GetAllFollowing gets all users that a given FID is following
// If fid is nil, it gets the following list for the authenticated user
func (w *Warpcast) GetAllFollowing(fid *int) (*UsersResult, error) {
	var users []ApiUser
	var cursor *string
	limit := 100

	// If fid is nil, get the authenticated user's FID
	if fid == nil {
		me, err := w.GetMe()
		if err != nil {
			return nil, fmt.Errorf("failed to get authenticated user: %w", err)
		}
		fid = &me.Fid
	}

	for {
		params := map[string]string{
			"fid":   fmt.Sprintf("%d", *fid),
			"limit": fmt.Sprintf("%d", limit),
		}
		if cursor != nil {
			params["cursor"] = *cursor
		}

		resp, err := w.request("GET", "following", params, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get following: %w", err)
		}

		var response FollowingGetResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal following response: %w", err)
		}

		if len(response.Result.Users) > 0 {
			users = append(users, response.Result.Users...)
		}

		if response.Next == nil {
			break
		}
		cursor = &response.Next.Cursor
	}

	return &UsersResult{Users: users}, nil
} 