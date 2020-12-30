package service

import (
	"encoding/json"
	"log"

	"github.com/thetinygoat/localeai/pkg/entities/ride"
)

// Service implements ride.Service
type Service struct {
	repo ride.Repository
}

// New instantiates a new ride service
func New(repo ride.Repository) ride.Service {
	return Service{repo: repo}
}

// Dump dumps data into the repository
func (s Service) Dump(data string) {
	ride := ride.Ride{}
	json.Unmarshal([]byte(data), &ride)
	err := s.repo.CreateRide(&ride)
	if err != nil {
		log.Fatal(err)
	}
}
