package farcaster

import (
	"encoding/json"
	"net/http"
)

type MeGetResponse struct {
	Result struct {
		User ApiUser `json:"user"`
	} `json:"result"`
}

func (w *Warpcast) getMe() (*ApiUser, error) {
	req, err := http.NewRequest("GET", w.config.BasePath+"me", nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MeGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	w.config.Username = &response.Result.User.Username
	return &response.Result.User, nil
}
