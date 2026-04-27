# Oaktioneer - Specification

## Overview
Fast GSP (Generalized Second Price) auction system for sponsored restaurant campaigns. Designed for sub-100ms latency.

## Auction Mechanism

### GSP (Generalized Second Price)
- **Ranking**: Restaurants ranked by `rankScore = qualityScore × bidAmount`
- **Pricing**: Winner pays `min(secondHighestBid, floorPrice)` or rank-adjusted minimum
- **Properties**: Strategy-proof, prevents jumping, incentivizes truthful bidding

### Quality Score (0.0 - 1.0)
Derived from:
- Order volume (weighted)
- Rating score
- Organic position metrics

### Input Parameters
- `numSlots` (X): Number of spots to award (1, 10, 40, configurable per request)
- `floorPrice`: Minimum bid threshold

## Data Models

### Restaurant
```
ID: string
Name: string
QualityScore: float64 (0.0 - 1.0)
BidAmount: float64
AutoBid: bool
MaxAutoBid: float64 (if auto-bidding enabled)
CampaignID: string (optional)
```

### Campaign
```
ID: string
RestaurantID: string
BidStrategy: "manual" | "auto"
MaxBudget: float64
TargetSlots: int
CreatedAt: timestamp
```

### AuctionRequest
```
NumSlots: int
FloorPrice: float64
FilterCriteria: optional
```

### AuctionResult
```
Winners: []WinnerPosition
ProcessingTimeMs: float64
Timestamp: int64
```

### WinnerPosition
```
Rank: int
Restaurant: Restaurant
CPM: float64 (cost per mille / display)
PayPrice: float64
```

## API Endpoints

### POST /auction
Run auction for X slots.

### POST /restaurants
Register/update a restaurant.

### POST /campaigns
Create auto-bidding campaign.

### GET /restaurants/:id
Get restaurant details.

### PUT /restaurants/:id/bid
Update bid amount.

## Performance Targets
- Single auction call: < 100ms P99
- Support: 1000+ restaurants per auction
- X (slots) configurable: 1-100

## Auto-Bidding Strategy
When enabled, system auto-bids based on:
- Target rank position
- Max budget constraints
- Quality score - higher quality = less needed to bid

## Implementation Notes
- Go for performance
- In-memory data store
-排序算法: sort.Slice with custom comparator for O(n log n)
- Pre-computed quality scores cached in memory