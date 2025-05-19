package model

import "time"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	DurationMin  int      `json:"duration_min"`
	Slots        []Slot   `json:"slots"`
	Participants []string `json:"participants"`
}

type Slot struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type Availability struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Slots   []Slot `json:"slots"`
}

type SlotSuggestion struct {
	Slot             Slot     `json:"slot"`
	UnavailableUsers []string `json:"unavailable_users"`
}
