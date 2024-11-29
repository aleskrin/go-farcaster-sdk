package farcaster

import (
	"fmt"
)

type IterableCastsResult struct {
	Casts  []ApiCast `json:"casts"`
	Cursor *string   `json:"cursor,omitempty"`
}

// GetCasts retrieves casts for a given FID (Farcaster ID) of a user
// fid: Farcaster ID of the user
// limit: Number of casts to retrieve (default 25, max 100)
func (c *Client) GetCasts(fid int, cursor *string, limit int) (*IterableCastsResult, error) {
	if limit <= 0 {
		limit = 25
	}
	if limit > 100 {
		limit = 100
	}

	var casts []ApiCast
	currentCursor := cursor

	for {
		params := map[string]interface{}{
			"fid":   fid,
			"limit": limit,
		}
		if currentCursor != nil {
			params["cursor"] = *currentCursor
		}

		response := &CastsGetResponse{}
		err := c.get("casts", params, response)
		if err != nil {
			return nil, fmt.Errorf("failed to get casts: %w", err)
		}

		if response.Result.Casts != nil {
			casts = append(casts, response.Result.Casts...)
		}

		// Break if we've reached the limit
		if response.Next == nil || len(casts) >= limit {
			break
		}
		currentCursor = response.Next.Cursor
	}

	// Trim casts to respect the limit
	if len(casts) > limit {
		casts = casts[:limit]
	}

	var nextCursor *string
	if currentCursor != nil {
		nextCursor = currentCursor
	}

	return &IterableCastsResult{
		Casts:  casts,
		Cursor: nextCursor,
	}, nil
}

