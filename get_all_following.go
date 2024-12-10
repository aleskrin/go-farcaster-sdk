package farcaster

import (
	"encoding/json"
	"fmt"
)

// GetAllFollowing retrieves all following for a user
// Parameters:
//   - fid: The FID of the user to retrieve following for
//
// Returns:
//   - *UsersResult: A collection of users
//   - error: Any error that occurred
func (w *Warpcast) GetAllFollowing(fid *int) (*UsersResult, error) {
	// If fid is nil, use authenticated user's fid
	userFid := fid
	if userFid == nil {
		me, err := w.getMe()
		if err != nil {
			return nil, err
		}
		userFid = &me.FID
	}

	var users []ApiUser
	var cursor *string
	limit := 100

	for {
		params := map[string]interface{}{
			"fid":    *userFid,
			"limit":  limit,
			"cursor": cursor,
		}

		response, err := w.request("GET", "following", nil, params, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get following: %w", err)
		}

		var followingResponse FollowingGetResponse
		if err := json.Unmarshal(response, &followingResponse); err != nil {
			return nil, err
		}

		if len(followingResponse.Result.Users) > 0 {
			users = append(users, followingResponse.Result.Users...)
		}

		if followingResponse.Next == nil {
			break
		}
		cursor = &followingResponse.Next.Cursor
	}

	return &UsersResult{Users: users}, nil
}
