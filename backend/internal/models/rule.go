package models

type Rule struct {
	ID         string   `json:"id"`
	GameID     string   `json:"game_id"`
	Countries  []string `json:"countries"`
	MinVersion string   `json:"min_version"`
	Platforms  []string `json:"platforms"`
	AppTypes   []string `json:"app_types"`
	IsActive   bool     `json:"is_active"`
}

type AccessRequest struct {
	GameID     string `json:"game_id"`
	Country    string `json:"country"`
	AppVersion string `json:"app_version"`
	Platform   string `json:"platform"`
	AppType    string `json:"app_type"`
}

type AccessResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason,omitempty"`
}
