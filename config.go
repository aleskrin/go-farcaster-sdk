package main

import (
	"fmt"
)

const FARCASTER_API_BASE_URL = "https://api.warpcast.com/v2/"

// ConfigurationParams holds the configuration parameters.
type ConfigurationParams struct {
	Username   *string            `json:"username,omitempty"`
	Password   *string            `json:"password,omitempty"`
	BasePath   string             `json:"base_path"`
	BaseOptions map[string]interface{} `json:"base_options,omitempty"`
}

// Configuration holds the configuration.
type Configuration struct {
	Params *ConfigurationParams `json:"params,omitempty"`
}

// NewConfiguration creates a new Configuration instance.
func NewConfiguration(data map[string]interface{}) (*Configuration, error) {
	params := &ConfigurationParams{
		BasePath: FARCASTER_API_BASE_URL,
	}

	if username, ok := data["username"].(string); ok {
		params.Username = &username
	}
	if password, ok := data["password"].(string); ok {
		params.Password = &password
	}
	if baseOptions, ok := data["base_options"].(map[string]interface{}); ok {
		params.BaseOptions = baseOptions
	}

	return &Configuration{
		Params: params,
	}, nil
}

func main() {
	// Example usage
	data := map[string]interface{}{
		"username": "user123",
		"password": "password123",
	}

	config, err := NewConfiguration(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Configuration:", config)
}
