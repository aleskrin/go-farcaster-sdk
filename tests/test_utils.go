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

func mockGetRecentCasts(cursor *string, limit int) []ApiCast {
	return []ApiCast{
		{
			Hash:       "0x1",
			ThreadHash: "0x2",
			ParentHash: "0x3",
			Author: Author{
				FID:          1,
				Username:     "1",
				DisplayName:  "1",
				RegisteredAt: nil,
				Pfp:          Pfp{URL: "https://lh3.googleusercontent.com/1", Verified: true},
				Profile:      Profile{Bio: Bio{Text: "1", Mentions: []string{}}},
				FollowerCount:  1,
				FollowingCount: 1,
				ReferrerUsername: nil,
				ViewerContext:    nil,
			},
			Text:          "1",
			Timestamp:     1675301079335,
			Mentions:      nil,
			Attachments:   nil,
			Embeds:        nil,
			Ancestors:     nil,
			Replies:       Replies{Count: 0},
			Reactions:     Reactions{Count: 0},
			Recasts:       Recasts{Count: 0, Recasters: []string{}},
			Watches:       Watches{Count: 0},
			Deleted:       nil,
			Recast:        nil,
			ViewerContext: nil,
		},
		{
			Hash:       "0x2",
			ThreadHash: "0x3",
			ParentHash: "0x4",
			Author: Author{
				FID:          2,
				Username:     "2",
				DisplayName:  "2",
				RegisteredAt: nil,
				Pfp:          Pfp{URL: "https://lh3.googleusercontent.com/2", Verified: true},
				Profile:      Profile{Bio: Bio{Text: "2", Mentions: []string{}}},
				FollowerCount:  2,
				FollowingCount: 2,
				ReferrerUsername: nil,
				ViewerContext:    nil,
			},
			Text:          "2",
			Timestamp:     1675301079335,
			Mentions:      nil,
			Attachments:   nil,
			Embeds:        nil,
			Ancestors:     nil,
			Replies:       Replies{Count: 0},
			Reactions:     Reactions{Count: 0},
			Recasts:       Recasts{Count: 0, Recasters: []string{}},
			Watches:       Watches{Count: 0},
			Deleted:       nil,
			Recast:        nil,
			ViewerContext: nil,
		},
		{
			Hash:       "0x3",
			ThreadHash: "0x4",
			ParentHash: "0x5",
			Author: Author{
				FID:          3,
				Username:     "3",
				DisplayName:  "3",
				RegisteredAt: nil,
				Pfp:          Pfp{URL: "https://lh3.googleusercontent.com/3", Verified: true},
				Profile:      Profile{Bio: Bio{Text: "3", Mentions: []string{}}},
				FollowerCount:  3,
				FollowingCount: 3,
				ReferrerUsername: nil,
				ViewerContext:    nil,
			},
			Text:          "3",
			Timestamp:     1675301079335,
			Mentions:      nil,
			Attachments:   nil,
			Embeds:        nil,
			Ancestors:     nil,
			Replies:       Replies{Count: 0},
			Reactions:     Reactions{Count: 0},
			Recasts:       Recasts{Count: 0, Recasters: []string{}},
			Watches:       Watches{Count: 0},
			Deleted:       nil,
			Recast:        nil,
			ViewerContext: nil,
		},
	}
}

func main() {
	users := MockGetRecentUsers(nil, 3)
	for _, user := range users {
		fmt.Printf("User: %+v\n", user)
	}
}
