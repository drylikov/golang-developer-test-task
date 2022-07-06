package loader

import (
	"app/provider"
	"app/storage"
	"log"
)

type Loader struct {
	Storage  storage.Storage
	Provider provider.Provider
}

func (l *Loader) Run() {
	var data = l.Provider.Load()
	l.Storage.FlushAll()
	l.Storage.Insert(data)

	log.Printf("Load data: %d parking taxi", len(data))
}
