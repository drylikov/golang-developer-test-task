package searchEngine

import (
	"app/models"
	"app/storage"
)

type SearchEngine struct {
	Storage storage.Storage
}

// https://medium.com/@ishagirdhar/import-cycles-in-golang-b467f9f0c5a0
type RequestConverter interface {
	ToQueryString() []string
}

func (s SearchEngine) Find(request RequestConverter) []models.ParkingTaxi {
	// We get go from our indices and get data on them.
	// You can save and write a lua script that will receive the ID and immediately get the data.
	return s.Storage.FindByQuery(request.ToQueryString())
}
