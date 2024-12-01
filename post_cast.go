package farcaster

import (
	"encoding/json"
	"fmt"
)

// PostCast posts a new cast to Farcaster
//
// Parameters:
//   - text: text of the cast
//   - embeds: optional list of embeds
//   - parent: optional parent of the cast
//   - channelKey: optional channel of the cast
//
// Returns:
//   - *CastContent: The result of posting the cast
//   - error: Any error that occurred
func (w *Warpcast) PostCast(text string, embeds []string, parent *Parent, channelKey *string) (*CastContent, error) {
	// Create request body
	body := CastsPostRequest{
		Text:       text,
		Embeds:     embeds,
		Parent:     parent,
		ChannelKey: channelKey,
	}

	// Make the request
	resp, err := w.request("POST", "casts", nil, body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to post cast: %w", err)
	}

	// Parse response
	var result CastsPostResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cast response: %w", err)
	}

	return &result.Result, nil
} 