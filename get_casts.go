package farcaster

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type IterableCastsResult struct {
	Casts  []ApiCast `json:"casts"`
	Cursor *string   `json:"cursor,omitempty"`
}

type ApiCast struct {
	Hash       string   `json:"hash"`
	ThreadHash string   `json:"threadHash"`
	ParentHash *string  `json:"parentHash,omitempty"`
	Author     ApiUser  `json:"author"`
	Text       string   `json:"text"`
}

type ApiUser struct {
	Fid       int    `json:"fid"`
	Username  string `json:"username"`
	DisplayName string `json:"displayName"`
	PfpUrl    string `json:"pfpUrl"`
}

type CastsGetResponse struct {
	Result struct {
		Casts []ApiCast `json:"casts"`
	} `json:"result"`
	Next struct {
		Cursor *string `json:"cursor"`
	} `json:"next,omitempty"`
}

// GetCasts retrieves casts for a given FID (Farcaster ID) of a user
// fid: Farcaster ID of the user
// limit: Number of casts to retrieve (default 25, max 100)
func (c *Warpcast) GetCasts(fid int, cursor *string, limit int) (*IterableCastsResult, error) {
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
		if response.Next.Cursor == nil || len(casts) >= limit {
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

func (c *Warpcast) get(endpoint string, params map[string]interface{}, response interface{}) error {
	req, err := http.NewRequest("GET", c.BaseURL+"/"+endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, fmt.Sprint(v))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}

