import (
	"fmt"
)

// ApiUser represents a user from the API
type ApiUser struct {
	// Add necessary fields here
}

// UsersResult represents a collection of users
type UsersResult struct {
	Users []ApiUser `json:"users"`
}

// IterableUsersResult represents a paginated collection of users
type IterableUsersResult struct {
	Users  []ApiUser `json:"users"`
	Cursor *string   `json:"cursor,omitempty"`
}

// FollowersResponse represents the API response for followers
type FollowersResponse struct {
	Result struct {
		Users []ApiUser `json:"users"`
	} `json:"result"`
	Next *struct {
		Cursor string `json:"cursor"`
	} `json:"next,omitempty"`
}

// GetFollowers retrieves followers for a user with pagination
func (c *Client) GetFollowers(fid int, cursor *string, limit int) (*IterableUsersResult, error) {
	if limit <= 0 {
		limit = 25
	}
	if limit > 100 {
		limit = 100
	}

	users := make([]ApiUser, 0)
	currentCursor := cursor
	
	for {
		params := map[string]interface{}{
			"fid":    fid,
			"cursor": currentCursor,
			"limit":  limit,
		}
		
		var response FollowersResponse
		if err := c.get("followers", params, &response); err != nil {
			return nil, fmt.Errorf("failed to get followers: %w", err)
		}

		if len(response.Result.Users) > 0 {
			users = append(users, response.Result.Users...)
		}

		if response.Next == nil || len(users) >= limit {
			break
		}
		currentCursor = &response.Next.Cursor
	}

	// Trim users to respect limit
	if len(users) > limit {
		users = users[:limit]
	}

	var finalCursor *string
	if response.Next != nil {
		finalCursor = &response.Next.Cursor
	}

	return &IterableUsersResult{
		Users:  users,
		Cursor: finalCursor,
	}, nil
}
