package store

import (
	"sync"

	"oaktioneer/internal/models"
)

type RestaurantStore struct {
	mu          sync.RWMutex
	restaurants map[string]*models.Restaurant
	campaigns   map[string]*models.Campaign
}

func NewRestaurantStore() *RestaurantStore {
	return &RestaurantStore{
		restaurants: make(map[string]*models.Restaurant),
		campaigns:   make(map[string]*models.Campaign),
	}
}

func (s *RestaurantStore) Get(rID string) (*models.Restaurant, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.restaurants[rID]
	return r, ok
}

func (s *RestaurantStore) Set(r *models.Restaurant) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.restaurants[r.ID] = r
}

func (s *RestaurantStore) GetAll() []*models.Restaurant {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*models.Restaurant, 0, len(s.restaurants))
	for _, r := range s.restaurants {
		result = append(result, r)
	}
	return result
}

func (s *RestaurantStore) GetCampaign(cID string) (*models.Campaign, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.campaigns[cID]
	return c, ok
}

func (s *RestaurantStore) SetCampaign(c *models.Campaign) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.campaigns[c.ID] = c
}

func (s *RestaurantStore) GetByCampaign(campaignID string) *models.Restaurant {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, r := range s.restaurants {
		if r.CampaignID == campaignID {
			return r
		}
	}
	return nil
}

func (s *RestaurantStore) UpdateBid(rID string, bid float64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r, ok := s.restaurants[rID]; ok {
		r.BidAmount = bid
		return true
	}
	return false
}