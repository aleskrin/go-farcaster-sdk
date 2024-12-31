package tests

import (
	"testing"
	"github.com/warpcast/warpcast-go/pkg/farcaster"
)

func TestGetCasts(t *testing.T) {
	client := farcaster.NewWarpcast() // You'll need to implement this constructor

	tests := []struct {
		name    string
		fid     int
		cursor  *string
		limit   int
		wantLen int
		wantErr bool
	}{
		{
			name:    "Valid FID with default limit",
			fid:     1,
			cursor:  nil,
			limit:   0, // Should default to 25
			wantLen: 25,
			wantErr: false,
		},
		{
			name:    "Valid FID with custom limit",
			fid:     1,
			cursor:  nil,
			limit:   10,
			wantLen: 10,
			wantErr: false,
		},
		{
			name:    "Invalid FID",
			fid:     -1,
			cursor:  nil,
			limit:   25,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "Exceeding max limit",
			fid:     1,
			cursor:  nil,
			limit:   150, // Should be capped at 100
			wantLen: 100,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetCasts(tt.fid, tt.cursor, tt.limit)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCasts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			// Check result length
			if len(result.Casts) != tt.wantLen {
				t.Errorf("GetCasts() got %d casts, want %d", len(result.Casts), tt.wantLen)
			}

			// Validate cast structure
			if len(result.Casts) > 0 {
				firstCast := result.Casts[0]
				if firstCast.Hash == "" {
					t.Error("GetCasts() first cast has empty hash")
				}
				if firstCast.Author.Fid == 0 {
					t.Error("GetCasts() first cast has invalid author FID")
				}
				if firstCast.Text == "" {
					t.Error("GetCasts() first cast has empty text")
				}
			}
		})
	}
}

func TestGetCastsPagination(t *testing.T) {
	client := farcaster.NewWarpcast()
	fid := 1
	limit := 10

	// First request
	result1, err := client.GetCasts(fid, nil, limit)
	if err != nil {
		t.Fatalf("First GetCasts() failed: %v", err)
	}

	if len(result1.Casts) != limit {
		t.Errorf("Expected %d casts, got %d", limit, len(result1.Casts))
	}

	// If there's no cursor, we can't test pagination
	if result1.Cursor == nil {
		t.Skip("No cursor returned, skipping pagination test")
	}

	// Second request with cursor
	result2, err := client.GetCasts(fid, result1.Cursor, limit)
	if err != nil {
		t.Fatalf("Second GetCasts() failed: %v", err)
	}

	if len(result2.Casts) != limit {
		t.Errorf("Expected %d casts in second page, got %d", limit, len(result2.Casts))
	}

	// Verify we got different casts
	for i, cast1 := range result1.Casts {
		for j, cast2 := range result2.Casts {
			if cast1.Hash == cast2.Hash {
				t.Errorf("Found duplicate cast (page1[%d] == page2[%d]): %s", i, j, cast1.Hash)
			}
		}
	}
}
