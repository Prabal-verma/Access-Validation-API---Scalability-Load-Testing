package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yourorg/auth-service/internal/models"
)

type RuleCache struct {
	client *redis.Client
}

func NewRuleCache(redisURL string) (*RuleCache, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	client := redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RuleCache{
		client: client,
	}, nil
}

func (c *RuleCache) GetRules(gameID string) []models.Rule {
	ctx := context.Background()
	key := fmt.Sprintf("rules:%s", gameID)

	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil
	}

	var rules []models.Rule
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil
	}

	return rules
}

func (c *RuleCache) SetRules(gameID string, rules []models.Rule) error {
	ctx := context.Background()
	key := fmt.Sprintf("rules:%s", gameID)

	data, err := json.Marshal(rules)
	if err != nil {
		return fmt.Errorf("failed to marshal rules: %w", err)
	}

	if err := c.client.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to set rules in cache: %w", err)
	}

	return nil
}
