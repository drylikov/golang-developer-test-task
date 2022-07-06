package fileProvider

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
)

type FileProvider struct {
	Path string
}

func (s FileProvider) Load() []models.ParkingTaxi {
	d, err := ioutil.ReadFile(s.Path)
	if err != nil {
		panic("File not found")
	}
	var result []models.ParkingTaxi
	if err := json.Unmarshal(d, &result); err != nil {
		panic("unable to unmarshal data")
	}
	return result
}
