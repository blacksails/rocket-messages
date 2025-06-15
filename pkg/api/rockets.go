package api

import "time"

type Rocket struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Speed           int       `json:"speed"`
	Mission         string    `json:"mission,omitempty"`
	Exploded        bool      `json:"exploded,omitempty"`
	ExplosionReason string    `json:"explosionReason,omitempty"`
	LastMessage     time.Time `json:"lastMessage"`
}

type ListRocketsRequest struct {
	SortDecending   bool            `json:"sortDecending,omitempty"`
	SortBy          []RocketSorting `json:"sortBy,omitempty"`
	IncludeExploded bool            `json:"includeExploded,omitempty"`
}

type RocketSorting string

const (
	RocketSortingType     RocketSorting = "type"
	RocketSortingSpeed    RocketSorting = "speed"
	RocketSortingMission  RocketSorting = "mission"
	RocketSortingExploded RocketSorting = "exploded"
)
