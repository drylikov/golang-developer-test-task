package server

import (
	"app/config"
	searchEngineService "app/services/searchEngine"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

var rnd *renderer.Render
var searchEngine *searchEngineService.SearchEngine

func Run(config config.WebServer, engine *searchEngineService.SearchEngine) {
	opts := renderer.Options{
		ParseGlobPattern: config.Tpl + config.TplPattern,
	}

	rnd = renderer.New(opts)
	searchEngine = engine

	fs := http.FileServer(http.Dir(config.StaticPath))
	mux := http.NewServeMux()

	setupRoutes(mux)
	mux.Handle(config.StaticPrefix, http.StripPrefix(config.StaticPrefix, fs))

	log.Println("Listening on port ", config.Port)
	err := http.ListenAndServe(config.Port, mux)

	if err != nil {
		panic(err)
	}
}
