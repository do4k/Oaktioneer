package main

import (
	"fmt"
	"time"

	"auction/internal/auction"
	"auction/internal/models"
	"auction/internal/store"
)

func main() {
	s := store.NewRestaurantStore()
	a := auction.New(s)

	for i := 0; i < 1000; i++ {
		s.Set(&models.Restaurant{
			ID:           fmt.Sprintf("rest-%d", i),
			Name:         fmt.Sprintf("Restaurant %d", i),
			QualityScore: float64(i%100) / 100.0,
			BidAmount:    float64(i%50) * 0.1,
		})
	}

	req := models.AuctionRequest{
		NumSlots:   10,
		FloorPrice: 0.0,
	}

	iterations := 10000
	var totalMs float64
	minMs := 1000000.0
	maxMs := 0.0

	for i := 0; i < iterations; i++ {
		start := time.Now()
		result := a.RunAuction(req)
		elapsed := float64(time.Since(start).Microseconds()) / 1000.0

		totalMs += elapsed
		if elapsed < minMs {
			minMs = elapsed
		}
		if elapsed > maxMs {
			maxMs = elapsed
		}

		if i == 0 {
			fmt.Printf("First run: %.2fms, winners: %d\n", elapsed, len(result.Winners))
		}
	}

	avgMs := totalMs / float64(iterations)

	fmt.Printf("=== Benchmark Results (%d restaurants, %d slots) ===\n", 1000, req.NumSlots)
	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Avg: %.2fms\n", avgMs)
	fmt.Printf("Min: %.2fms\n", minMs)
	fmt.Printf("Max: %.2fms\n", maxMs)
	fmt.Printf("\nTarget: <100ms\n")
	if avgMs < 100 {
		fmt.Printf("PASS: Average %.2fms is under 100ms\n", avgMs)
	} else {
		fmt.Printf("FAIL: Average %.2fms exceeds 100ms\n", avgMs)
	}
}