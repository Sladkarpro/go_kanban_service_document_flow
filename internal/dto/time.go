package dto

import (
	"encoding/json"
	"errors"
	"time"
)

var errEstimateRequired = errors.New("estimateMinutes required")

type SetEstimateRequest struct {
	EstimateMinutes *int `json:"estimateMinutes"`
}

func (r *SetEstimateRequest) UnmarshalJSON(data []byte) error {
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	raw, ok := payload["estimateMinutes"]
	if !ok {
		return errEstimateRequired
	}
	if string(raw) == "null" {
		return nil
	}
	var minutes int
	if err := json.Unmarshal(raw, &minutes); err != nil {
		return err
	}
	r.EstimateMinutes = &minutes
	return nil
}

type EstimateResponse struct {
	ID              int64     `json:"id"`
	EstimateMinutes *int      `json:"estimateMinutes"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateTimeEntryRequest struct {
	Minutes int        `json:"minutes"`
	WorkAt  *time.Time `json:"workAt,omitempty"`
}

type TimeEntryResponse struct {
	ID            int64     `json:"id"`
	Minutes       int       `json:"minutes"`
	WorkAt        time.Time `json:"workAt"`
	LoggedMinutes int       `json:"loggedMinutes"`
}

type DeleteTimeEntryResponse struct {
	LoggedMinutes int `json:"loggedMinutes"`
}
