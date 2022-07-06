package provider

import (
	"app/models"
	"app/provider/fileProvider"
	"app/provider/httpProvider"
	"log"
	"os"
)

type Provider interface {
	Load() []models.ParkingTaxi
}

// CreateProvider is a Factory Method
func StrategyFactoryProvider(path string) Provider {

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			log.Printf("Use HTTP strategy from url: %s", path)
			return &httpProvider.HttpProvider{Path: path}
		}
	}

	log.Printf("Use local file strategy path: %s", path)
	return &fileProvider.FileProvider{Path: path}
}
