package main

import (
	"log"
	"net/http"
	"os"

	"oaktioneer/internal/api"
	"oaktioneer/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s := store.NewRestaurantStore()
	h := api.New(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/restaurants", h.RegisterRestaurant)
	mux.HandleFunc("/restaurants/", h.GetRestaurant)
	mux.HandleFunc("/restaurants/-/bid", h.UpdateBid)
	mux.HandleFunc("/auction", h.RunAuction)
	mux.HandleFunc("/campaigns", h.RegisterCampaign)

	log.Printf("Starting Oaktioneer on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}