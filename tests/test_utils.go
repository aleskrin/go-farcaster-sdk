package main

import "fmt"

type ProfileBio struct {
	Text     string   `json:"text"`
	Mentions []string `json:"mentions"`
}

type Profile struct {
	Bio ProfileBio `json:"bio"`
}

type PFP struct {
	URL      string `json:"url"`
	Verified bool   `json:"verified"`
}

type ViewerContext struct {
	Following           bool  `json:"following"`
	FollowedBy          bool  `json:"followed_by"`
	CanSendDirectCasts  *bool `json:"can_send_direct_casts"`
}

type ApiUser struct {
	FID             int           `json:"fid"`
	Username        string        `json:"username"`
	DisplayName     string        `json:"display_name"`
	RegisteredAt    *string       `json:"registered_at"`
	PFP             PFP           `json:"pfp"`
	Profile         Profile       `json:"profile"`
	FollowerCount   int           `json:"follower_count"`
	FollowingCount  int           `json:"following_count"`
	ReferrerUsername *string      `json:"referrer_username"`
	ViewerContext   ViewerContext `json:"viewer_context"`
}

func MockGetRecentUsers(cursor *string, limit int) []ApiUser {
	return []ApiUser{
		{
			FID:          1,
			Username:     "hello",
			DisplayName:  "world",
			RegisteredAt: nil,
			PFP:          PFP{URL: "https://openseauserdata.com/files/20.svg", Verified: true},
			Profile:      Profile{Bio: ProfileBio{Text: "foo", Mentions: []string{}}},
			FollowerCount:   1,
			FollowingCount:  2,
			ReferrerUsername: nil,
			ViewerContext: ViewerContext{
				Following:          false,
				FollowedBy:         false,
				CanSendDirectCasts: nil,
			},
		},
		{
			FID:          2,
			Username:     "hello1",
			DisplayName:  "world1",
			RegisteredAt: nil,
			PFP:          PFP{URL: "https://openseauserdata.com/files/20.svg", Verified: true},
			Profile:      Profile{Bio: ProfileBio{Text: "foo1", Mentions: []string{}}},
			FollowerCount:   1,
			FollowingCount:  2,
			ReferrerUsername: nil,
			ViewerContext: ViewerContext{
				Following:          false,
				FollowedBy:         false,
				CanSendDirectCasts: nil,
			},
		},
		{
			FID:          3,
			Username:     "hello2",
			DisplayName:  "world2",
			RegisteredAt: nil,
			PFP:          PFP{URL: "https://openseauserdata.com/files/20.svg", Verified: true},
			Profile:      Profile{Bio: ProfileBio{Text: "foo2", Mentions: []string{}}},
			FollowerCount:   1,
			FollowingCount:  2,
			ReferrerUsername: nil,
			ViewerContext: ViewerContext{
				Following:          false,
				FollowedBy:         false,
				CanSendDirectCasts: nil,
			},
		},
	}
}

func main() {
	users := MockGetRecentUsers(nil, 3)
	for _, user := range users {
		fmt.Printf("User: %+v\n", user)
	}
}
