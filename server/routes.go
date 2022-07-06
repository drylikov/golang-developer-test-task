package server

import (
	v1 "app/server/api/v1"
	"app/server/web"
	"net/http"
)

func setupRoutes(mux *http.ServeMux) {
	// Inject Dependencies
	web.Rnd = rnd
	v1.Rnd = rnd
	v1.SearchEngine = searchEngine

	mux.HandleFunc("/", web.Home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/api/v1/search/taxi/parking", v1.Search)
}

func about(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}
