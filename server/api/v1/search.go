package v1

import (
	"app/services/searchEngine"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"time"
)

var Rnd *renderer.Render
var SearchEngine *searchEngine.SearchEngine

type RequestConverter struct {
	GlobalID string
	ID       string
	Mode     string
}

func (r RequestConverter) ToQueryString() []string {
	var queryString []string
	if r.ID != "" {
		queryString = append(queryString, "ID:"+r.ID)
	}

	if r.Mode != "" {
		queryString = append(queryString, "Mode:"+r.Mode)
	}

	if r.GlobalID != "" {
		queryString = append(queryString, "GlobalID:"+r.GlobalID)
	}

	return queryString
}

func Search(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	keys := req.URL.Query()
	//time.Sleep(2 * time.Second)
	if len(keys) == 0 {
		errorMessage := "None request params"
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	request := &RequestConverter{GlobalID: keys.Get("global_id"), ID: keys.Get("id"), Mode: keys.Get("mode")}
	parkingTaxi := SearchEngine.Find(request)

	elapsed := time.Since(start)
	log.Printf("Search func: %s", elapsed)

	Rnd.JSON(w, http.StatusOK, parkingTaxi)
}
