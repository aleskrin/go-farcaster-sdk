package tests

import (
	"testing"
	"your_project_name/models" // adjust this import path based on your project structure
)

func TestGetCasts(t *testing.T) {
	tests := []struct {
		name    string
		movieID int
		wantLen int
		wantErr bool
	}{
		{
			name:    "Valid movie ID",
			movieID: 550, // Fight Club as example
			wantLen: 10,  // Expecting at least 10 cast members
			wantErr: false,
		},
		{
			name:    "Invalid movie ID",
			movieID: -1,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "Non-existent movie ID",
			movieID: 999999999,
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			casts, err := models.GetCasts(tt.movieID)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCasts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check result length for valid cases
			if !tt.wantErr && len(casts) < tt.wantLen {
				t.Errorf("GetCasts() got %d casts, want at least %d", len(casts), tt.wantLen)
			}

			// For valid cases, check that cast members have required fields
			if !tt.wantErr && len(casts) > 0 {
				firstCast := casts[0]
				if firstCast.Name == "" {
					t.Error("GetCasts() first cast member has empty name")
				}
				if firstCast.Character == "" {
					t.Error("GetCasts() first cast member has empty character")
				}
				if firstCast.ProfilePath == "" {
					t.Error("GetCasts() first cast member has empty profile path")
				}
			}
		})
	}
}

func TestGetCastsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Test with a known movie (Fight Club)
	movieID := 550
	casts, err := models.GetCasts(movieID)

	if err != nil {
		t.Fatalf("GetCasts() failed with error: %v", err)
	}

	// Verify we got some cast members
	if len(casts) == 0 {
		t.Error("GetCasts() returned empty cast list for Fight Club")
	}

	// Check for specific cast members we know should be in Fight Club
	foundBradPitt := false
	foundEdwardNorton := false

	for _, cast := range casts {
		switch cast.Name {
		case "Brad Pitt":
			foundBradPitt = true
			if cast.Character != "Tyler Durden" {
				t.Errorf("Expected Brad Pitt to play 'Tyler Durden', got '%s'", cast.Character)
			}
		case "Edward Norton":
			foundEdwardNorton = true
			if cast.Character != "The Narrator" {
				t.Errorf("Expected Edward Norton to play 'The Narrator', got '%s'", cast.Character)
			}
		}
	}

	if !foundBradPitt {
		t.Error("Brad Pitt not found in Fight Club cast")
	}
	if !foundEdwardNorton {
		t.Error("Edward Norton not found in Fight Club cast")
	}
}
