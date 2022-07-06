package storage

import (
	"app/config"
	"app/models"
	"app/storage/redis"
)

type Storage interface {
	Insert(data []models.ParkingTaxi)
	FlushAll()
	GetPatternName() string
	GetAllowedIndex() []string
	FindByQuery(query []string) []models.ParkingTaxi
}

func GetStorage(storage config.Storage) Storage {
	// sry, idk how to do it better
	return redis.Connect(storage.Redis)
}
