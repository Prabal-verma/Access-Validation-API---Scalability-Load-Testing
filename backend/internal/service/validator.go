package service

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/yourorg/auth-service/internal/models"
)

type RuleCache interface {
	GetRules(gameID string) []models.Rule
	SetRules(gameID string, rules []models.Rule) error
}

type Validator struct {
	cache RuleCache
}

func NewValidator(cache RuleCache) *Validator {
	return &Validator{
		cache: cache,
	}
}

func (v *Validator) ValidateAccess(req *models.AccessRequest) (*models.AccessResponse, error) {
	rules := v.cache.GetRules(req.GameID)
	if len(rules) == 0 {
		return &models.AccessResponse{
			Allowed: false,
			Reason:  "no rules found for game",
		}, nil
	}

	reqVersion, err := version.NewVersion(req.AppVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid app version: %w", err)
	}

	for _, rule := range rules {
		if !rule.IsActive {
			continue
		}

		// Check minimum version
		minVersion, err := version.NewVersion(rule.MinVersion)
		if err != nil {
			return nil, fmt.Errorf("invalid rule min version: %w", err)
		}

		if reqVersion.LessThan(minVersion) {
			continue
		}

		// Check country
		countryMatch := false
		for _, country := range rule.Countries {
			if strings.EqualFold(country, req.Country) {
				countryMatch = true
				break
			}
		}
		if !countryMatch {
			continue
		}

		// Check platform
		platformMatch := false
		for _, platform := range rule.Platforms {
			if strings.EqualFold(platform, req.Platform) {
				platformMatch = true
				break
			}
		}
		if !platformMatch {
			continue
		}

		// Check app type
		appTypeMatch := false
		for _, appType := range rule.AppTypes {
			if strings.EqualFold(appType, req.AppType) {
				appTypeMatch = true
				break
			}
		}
		if !appTypeMatch {
			continue
		}

		return &models.AccessResponse{
			Allowed: true,
		}, nil
	}

	return &models.AccessResponse{
		Allowed: false,
		Reason:  "no matching rules found",
	}, nil
}
