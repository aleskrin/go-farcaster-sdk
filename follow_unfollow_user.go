package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Define types for request and response
type FollowsPutRequest struct {
	TargetFID int `json:"target_fid"`
}

type FollowsDeleteRequest struct {
	TargetFID int `json:"target_fid"`
}

type StatusResponse struct {
	Result string `json:"result"`
}

// API Client structure
type APIClient struct {
	BaseURL string
	Client  *http.Client
}

// FollowUser sends a request to follow a user
func (c *APIClient) FollowUser(fid int) (string, error) {
	body := FollowsPutRequest{TargetFID: fid}
	bodyData, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, c.BaseURL+"/follows", bytes.NewBuffer(bodyData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return statusResp.Result, nil
}

// UnfollowUser sends a request to unfollow a user
func (c *APIClient) UnfollowUser(fid int) (string, error) {
	body := FollowsDeleteRequest{TargetFID: fid}
	bodyData, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodDelete, c.BaseURL+"/follows", bytes.NewBuffer(bodyData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return statusResp.Result, nil
}

func main() {
	client := &APIClient{
		BaseURL: "https://api.example.com",
		Client:  &http.Client{},
	}

	// Example usage of FollowUser and UnfollowUser
	followResult, err := client.FollowUser(123)
	if err != nil {
		fmt.Println("Error following user:", err)
	} else {
		fmt.Println("Follow result:", followResult)
	}

	unfollowResult, err := client.UnfollowUser(123)
	if err != nil {
		fmt.Println("Error unfollowing user:", err)
	} else {
		fmt.Println("Unfollow result:", unfollowResult)
	}
}

