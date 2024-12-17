package main

import (
	"log"
	"os"
	"testing"
	"github.com/joho/godotenv"
)

type Warpcast struct {
	AccessToken string
}

func NewWarpcast(accessToken string) *Warpcast {
	return &Warpcast{AccessToken: accessToken}
}

func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
}

func GetAccessToken() string {
	token := os.Getenv("AUTH")
	if token == "" {
		log.Fatal("AUTH environment variable not set")
	}
	return token
}

func TestMain(m *testing.M) {
	LoadEnvironment()
	os.Exit(m.Run())
}

func TestClientSetup(t *testing.T) {
	accessToken := GetAccessToken()
	client := NewWarpcast(accessToken)

	if client.AccessToken == "" {
		t.Fatalf("Expected a valid access token, got an empty string")
	}
}

func TestVCRConfig(t *testing.T) {
	vcrConfig := map[string]string{
		"filter_headers": "authorization,DUMMY",
	}

	if vcrConfig["filter_headers"] == "" {
		t.Fatalf("Expected filter_headers to be set, got an empty value")
	}
}
