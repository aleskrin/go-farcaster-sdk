package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Warpcast struct {
	AccessToken      string
	Wallet           string
	RotationDuration int
	ExpiresAt        int64
}

func NewWarpcastFromMnemonic(mnemonic string, rotationDuration int) *Warpcast {
	return &Warpcast{
		AccessToken:      "dummy_access_token",
		Wallet:           "dummy_wallet",
		RotationDuration: rotationDuration,
	}
}

func NewWarpcastFromPrivateKey(privateKey string) *Warpcast {
	return &Warpcast{
		AccessToken:      "dummy_access_token",
		Wallet:           "dummy_wallet",
		ExpiresAt:        time.Now().Add(10 * time.Minute).UnixMilli(),
	}
}

func NewWarpcastFromAccessToken(authToken string) *Warpcast {
	return &Warpcast{
		AccessToken: authToken,
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func clientFromMnemonic() {
	loadEnv()
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		panic("MNEMONIC env var not set")
	}
	client := NewWarpcastFromMnemonic(mnemonic, 200)
	fmt.Println(client.AccessToken)
	fmt.Println("client.GetMe():", "dummy_user")
}

func clientFromPrivateKey() {
	loadEnv()
	privateKey := os.Getenv("PKEY")
	if privateKey == "" {
		panic("PKEY env var not set")
	}
	client := NewWarpcastFromPrivateKey(privateKey)
	fmt.Println(client.ExpiresAt)
	fmt.Println(client.AccessToken)
}

func clientFromAuth() {
	loadEnv()
	authToken := os.Getenv("AUTH")
	if authToken == "" {
		panic("AUTH env var not set")
	}
	client := NewWarpcastFromAccessToken(authToken)
	fmt.Println("client.GetMe():", "dummy_user")
}

func testRotation() {
	loadEnv()
	privateKey := os.Getenv("PKEY")
	if privateKey == "" {
		panic("PKEY env var not set")
	}
	client := NewWarpcastFromPrivateKey(privateKey)
	for {
		fmt.Println(client.AccessToken)
		fmt.Println(client.ExpiresAt)
		time.Sleep(25 * time.Second)
	}
}

func testStreamCasts() {
	loadEnv()
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		panic("MNEMONIC env var not set")
	}
	client := NewWarpcastFromMnemonic(mnemonic, 200)
	fmt.Println(client.AccessToken)
	for i := 0; i < 5; i++ {
		fmt.Println("dummy_cast_hash")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	clientFromMnemonic()
	clientFromPrivateKey()
	clientFromAuth()
	testRotation()
	testStreamCasts()
}
