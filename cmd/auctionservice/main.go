package main

import (
	"log"
	"net/http"

	"oaktioneer/internal/api"
	"oaktioneer/internal/store"
)

func main() {
	s := store.NewRestaurantStore()
	h := api.New(s)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("POST /restaurants", h.RegisterRestaurant)
	mux.HandleFunc("GET /restaurants/{id}", h.GetRestaurant)
	mux.HandleFunc("PUT /restaurants/{id}/bid", h.UpdateBid)
	mux.HandleFunc("POST /auction", h.RunAuction)
	mux.HandleFunc("POST /campaigns", h.RegisterCampaign)

	log.Println("Starting auction service on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}