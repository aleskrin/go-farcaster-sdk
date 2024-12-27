package main

import (
	"errors"
	"fmt"
	"time"
)

type MockResponse struct{}

type MockResponsePut struct{}

type AuthParams struct {
	Timestamp int64 `json:"timestamp"`
	ExpiresAt int64 `json:"expiresAt"`
}

type Token struct {
	Secret    string `json:"secret"`
	ExpiresAt int64  `json:"expiresAt"`
}

type Result struct {
	Success string `json:"success"`
}

type TokenResponse struct {
	Token Token `json:"token"`
}

func (m MockResponse) JSON() map[string]interface{} {
	return map[string]interface{}{
		"result": map[string]string{
			"success": "true",
		},
	}
}

func (m MockResponsePut) JSON() map[string]interface{} {
	return map[string]interface{}{
		"result": map[string]interface{}{
			"token": map[string]interface{}{
				"secret":    "MK-ABC123...",
				"expiresAt": 1610000000000,
			},
		},
	}
}

func NowMs() int64 {
	return time.Now().UnixMilli()
}

type WarpcastClient struct{
	WalletSet bool
}

func (c *WarpcastClient) CreateNewAuthToken(expiresIn int) error {
	if !c.WalletSet {
		return errors.New("Wallet not set")
	}
	return nil
}

func (c *WarpcastClient) DeleteAuth() (Result, error) {
	mockResponse := MockResponse{}
	result := mockResponse.JSON()["result"].(map[string]string)
	return Result{Success: result["success"]}, nil
}

func (c *WarpcastClient) PutAuth(authParams AuthParams) (Token, error) {
	mockResponse := MockResponsePut{}
	response := mockResponse.JSON()["result"].(map[string]interface{})["token"].(map[string]interface{})
	return Token{
		Secret:    response["secret"].(string),
		ExpiresAt: int64(response["expiresAt"].(int)),
	}, nil
}

func main() {
	client := &WarpcastClient{WalletSet: false}

	// Test auth params
	now := time.Now().Unix()
	authParams := AuthParams{
		Timestamp: now * 1000,
		ExpiresAt: (now + 600) * 1000,
	}
	fmt.Printf("AuthParams: %+v\n", authParams)

	fmt.Printf("NowMs: %d\n", NowMs())

	if err := client.CreateNewAuthToken(10); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if result, err := client.DeleteAuth(); err == nil {
		fmt.Printf("DeleteAuth Result: %+v\n", result)
	} else {
		fmt.Printf("Error: %s\n", err)
	}

	if token, err := client.PutAuth(authParams); err == nil {
		fmt.Printf("PutAuth Token: %+v\n", token)
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}

