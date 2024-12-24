package main

import (
	"errors"
	"testing"
)

type Warpcast struct {}

func (w *Warpcast) GetBasePath() bool {
	return true
}

func (w *Warpcast) GetBaseOptions() interface{} {
	return nil
}

func GetWallet(options ...map[string]string) (interface{}, error) {
	if len(options) > 0 {
		if mnemonic, exists := options[0]["mnemonic"]; exists && mnemonic != "" {
			return nil, errors.New("mnemonic option provided")
		}
		if privateKey, exists := options[0]["private_key"]; exists && privateKey != "" {
			return nil, errors.New("private key option provided")
		}
	}
	return nil, nil
}

func TestGetBasePath(t *testing.T) {
	client := &Warpcast{}
	if !client.GetBasePath() {
		t.Errorf("expected true, got false")
	}
}

func TestGetBaseOptions(t *testing.T) {
	client := &Warpcast{}
	if client.GetBaseOptions() != nil {
		t.Errorf("expected nil, got %v", client.GetBaseOptions())
	}
}

func TestGetWallet(t *testing.T) {
	if wallet, err := GetWallet(); err != nil || wallet != nil {
		t.Errorf("expected nil wallet and no error, got wallet: %v, error: %v", wallet, err)
	}

	_, err := GetWallet(map[string]string{"mnemonic": "test"})
	if err == nil {
		t.Errorf("expected error for mnemonic, got nil")
	}

	_, err = GetWallet(map[string]string{"private_key": "test"})
	if err == nil {
		t.Errorf("expected error for private_key, got nil")
	}
}
