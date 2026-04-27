# Restaurant Auction System

Fast GSP (Generalized Second Price) auction system for sponsored restaurant campaigns. Designed for sub-100ms latency.

## Quick Start

```bash
go run ./cmd/auctionservice
```

Server runs on `localhost:8080`.

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/restaurants` | POST | Register a restaurant |
| `/restaurants/{id}` | GET | Get restaurant |
| `/restaurants/{id}/bid` | PUT | Update bid |
| `/auction` | POST | Run auction for X slots |
| `/campaigns` | POST | Create auto-bid campaign |
| `/health` | GET | Health check |

## Auction Mechanism

**GSP (Generalized Second Price)**:
- Ranking: `rankScore = qualityScore × bidAmount`
- Winner pays second-highest qualified bid (or floor)

## Examples

```bash
# Register restaurant
curl -X POST localhost:8080/restaurants \
  -H "Content-Type: application/json" \
  -d '{"id":"r1","name":"Pizza Place","qualityScore":0.95,"bidAmount":1.50}'

# Run auction (top 10)
curl -X POST localhost:8080/auction \
  -H "Content-Type: application/json" \
  -d '{"numSlots":10,"floorPrice":0.0}'

# Auto-bid campaign
curl -X POST localhost:8080/campaigns \
  -H "Content-Type: application/json" \
  -d '{"id":"c1","restaurantId":"r1","bidStrategy":"auto","maxBudget":10.0,"targetSlots":3}'
```

## Benchmark

```
=== Benchmark Results (1000 restaurants, 10 slots) ===
Iterations: 10000
Avg: 0.09ms
Min: 0.05ms
Max: 0.42ms

Target: <100ms
PASS: Average is under 100ms
```

Run benchmark: `go run ./cmd/benchmark`