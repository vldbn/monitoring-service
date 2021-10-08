package schedule

import (
	"log"
	"monitoring-service/internal/app/service"
	"time"
)

// Schedule struct for running periodical tasks
type Schedule struct {
	services service.Service
}

// TaskUpdateCryptocurrenciesRates method for updating Cryptocurrency rates
func (s *Schedule) TaskUpdateCryptocurrenciesRates() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		log.Println("Starting task - Update Rates")
		if err := s.services.Cryptocurrency().UpdateCryptocurrencies(); err != nil {
			log.Printf("Schedule task error: %s", err.Error())
			return
		}
		log.Println("Finished task - Updated rates")
	}
	time.Sleep(10 * time.Second)
	ticker.Stop()
}

// NewSchedule constructor
func NewSchedule(services service.Service) *Schedule {
	return &Schedule{services: services}
}
