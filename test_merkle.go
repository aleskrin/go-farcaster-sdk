package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Warpcast struct{}

func (c *Warpcast) GetMentionAndReplyNotifications(limit int) *NotificationResponse {
	return &NotificationResponse{Notifications: make([]string, 150)}
}

func (c *Warpcast) GetUserCollections(ownerFID, limit int) *CollectionResponse {
	return &CollectionResponse{Collections: make([]string, 102)}
}

func (c *Warpcast) GetCollectionOwners(collectionID string, limit int) *OwnerResponse {
	return &OwnerResponse{Users: make([]string, 102)}
}

func (c *Warpcast) GetHealthcheck() *HealthcheckResponse {
	return &HealthcheckResponse{}
}

type NotificationResponse struct {
	Notifications []string
}

type CollectionResponse struct {
	Collections []string
}

type OwnerResponse struct {
	Users []string
}

type HealthcheckResponse struct{}

func TestMentionReplyNotifications(t *testing.T) {
	client := &Warpcast{}
	response := client.GetMentionAndReplyNotifications(150)

	assert.NotNil(t, response)
	assert.GreaterOrEqual(t, len(response.Notifications), 1, "Expected at least one notification")
}

func TestGetUserCollections(t *testing.T) {
	client := &Warpcast{}
	response := client.GetUserCollections(2, 101)

	assert.NotNil(t, response)
	assert.Greater(t, len(response.Collections), 1, "Expected more than one collection")
}

func TestGetCollectionOwners(t *testing.T) {
	client := &Warpcast{}
	response := client.GetCollectionOwners("proof-of-merge", 10000)

	assert.NotNil(t, response)
	assert.Greater(t, len(response.Users), 101, "Expected more than 101 users")
}

func TestGetHealthcheck(t *testing.T) {
	client := &Warpcast{}
	response := client.GetHealthcheck()

	assert.NotNil(t, response, "Expected healthcheck response to be non-nil")
}

