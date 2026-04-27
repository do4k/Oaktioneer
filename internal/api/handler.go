package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"oaktioneer/internal/auction"
	"oaktioneer/internal/models"
	"oaktioneer/internal/store"
)

type Handler struct {
	store   *store.RestaurantStore
	auctor  *auction.Auctor
}

func New(store *store.RestaurantStore) *Handler {
	return &Handler{
		store:  store,
		auctor: auction.New(store),
	}
}

func (h *Handler) RegisterRestaurant(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.ListRestaurants(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.UpdateBid(w, r)
		return
	}

	var req models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.store.Set(&req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func (h *Handler) ListRestaurants(w http.ResponseWriter, r *http.Request) {
	id := getPathID(r)
	if id != "" {
		h.GetRestaurant(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.store.GetAll())
}

func (h *Handler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	id := getPathID(r)
	restaurant, ok := h.store.Get(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}

func (h *Handler) UpdateBid(w http.ResponseWriter, r *http.Request) {
	id := getPathID(r)

	var req struct {
		BidAmount float64 `json:"bidAmount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !h.store.UpdateBid(id, req.BidAmount) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	restaurant, _ := h.store.Get(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}

func (h *Handler) RunAuction(w http.ResponseWriter, r *http.Request) {
	var req models.AuctionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.auctor.RunAuction(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) RegisterCampaign(w http.ResponseWriter, r *http.Request) {
	var req models.Campaign
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.CreatedAt = time.Now()
	h.store.SetCampaign(&req)

	if r, ok := h.store.Get(req.RestaurantID); ok {
		r.CampaignID = req.ID
		r.AutoBid = req.BidStrategy == "auto"
		r.MaxAutoBid = req.MaxBudget
		h.store.Set(r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func getPathID(r *http.Request) string {
	path := r.URL.Path
	idx := strings.LastIndex(path, "/")
	if idx > 0 {
		return path[idx+1:]
	}
	return ""
}