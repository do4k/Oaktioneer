package models

import "time"

type Restaurant struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	QualityScore float64 `json:"qualityScore"`
	BidAmount    float64 `json:"bidAmount"`
	AutoBid      bool    `json:"autoBid"`
	MaxAutoBid   float64 `json:"maxAutoBid,omitempty"`
	CampaignID  string  `json:"campaignId,omitempty"`
}

type Campaign struct {
	ID           string    `json:"id"`
	RestaurantID string    `json:"restaurantId"`
	BidStrategy  string    `json:"bidStrategy"` // "manual" | "auto"
	MaxBudget    float64   `json:"maxBudget"`
	TargetSlots  int       `json:"targetSlots"`
	CreatedAt   time.Time `json:"createdAt"`
}

type AuctionRequest struct {
	NumSlots      int     `json:"numSlots"`
	FloorPrice    float64 `json:"floorPrice"`
	FilterCriteria string `json:"filterCriteria,omitempty"`
}

type WinnerPosition struct {
	Rank        int       `json:"rank"`
	Restaurant  Restaurant `json:"restaurant"`
	CPM         float64   `json:"cpm"`
	PayPrice    float64   `json:"payPrice"`
}

type AuctionResult struct {
	Winners        []WinnerPosition `json:"winners"`
	ProcessingTimeMs float64         `json:"processingTimeMs"`
	Timestamp      int64           `json:"timestamp"`
}