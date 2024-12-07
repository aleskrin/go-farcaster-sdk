package main

import (
	"encoding/json"
)

type Client struct {
	// Add relevant fields and methods for the Client struct
}

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

type ApiUser struct {
	// Add relevant fields based on your API response
}

func (c *Client) GetAllFollowing(fid *int) (*UsersResult, error) {
	// If fid is nil, use authenticated user's fid
	userFid := fid
	if userFid == nil {
		me, err := c.GetMe()
		if err != nil {
			return nil, err
		}
		userFid = &me.Fid
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

		response, err := c.get("following", params)
		if err != nil {
			return nil, err
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
