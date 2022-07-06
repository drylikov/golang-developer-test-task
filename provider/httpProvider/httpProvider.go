package httpProvider

import (
	"app/models"
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpProvider struct {
	Path string
}

func (s HttpProvider) Load() []models.ParkingTaxi {
	resp, err := http.Get(s.Path)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var ParkingTaxi []models.ParkingTaxi
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &ParkingTaxi); err != nil {
		panic(err)
	}

	return ParkingTaxi
}
