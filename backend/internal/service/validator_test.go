package service

import (
	"testing"

	"github.com/yourorg/auth-service/internal/models"
)

type mockCache struct {
	rules map[string][]models.Rule
}

func (m *mockCache) GetRules(gameID string) []models.Rule {
	return m.rules[gameID]
}

func (m *mockCache) SetRules(gameID string, rules []models.Rule) error {
	m.rules[gameID] = rules
	return nil
}

func TestValidateAccess(t *testing.T) {
	cache := &mockCache{
		rules: map[string][]models.Rule{
			"game1": {
				{
					ID:         "rule1",
					GameID:     "game1",
					Countries:  []string{"US", "CA"},
					MinVersion: "1.0.0",
					Platforms:  []string{"iOS", "Android"},
					AppTypes:   []string{"mobile"},
					IsActive:   true,
				},
			},
		},
	}

	validator := NewValidator(cache)

	tests := []struct {
		name    string
		req     *models.AccessRequest
		want    bool
		wantErr bool
	}{
		{
			name: "valid request",
			req: &models.AccessRequest{
				GameID:     "game1",
				Country:    "US",
				AppVersion: "1.0.0",
				Platform:   "iOS",
				AppType:    "mobile",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "invalid version",
			req: &models.AccessRequest{
				GameID:     "game1",
				Country:    "US",
				AppVersion: "0.9.0",
				Platform:   "iOS",
				AppType:    "mobile",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "invalid country",
			req: &models.AccessRequest{
				GameID:     "game1",
				Country:    "UK",
				AppVersion: "1.0.0",
				Platform:   "iOS",
				AppType:    "mobile",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "invalid platform",
			req: &models.AccessRequest{
				GameID:     "game1",
				Country:    "US",
				AppVersion: "1.0.0",
				Platform:   "Windows",
				AppType:    "mobile",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "invalid app type",
			req: &models.AccessRequest{
				GameID:     "game1",
				Country:    "US",
				AppVersion: "1.0.0",
				Platform:   "iOS",
				AppType:    "desktop",
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validator.ValidateAccess(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Allowed != tt.want {
				t.Errorf("ValidateAccess() = %v, want %v", got.Allowed, tt.want)
			}
		})
	}
}
