package main

import (
	"app/cmd"
	configLoader "app/config"
	"app/provider"
	"app/server"
	dataLoaderService "app/services/loader"
	searchEngineService "app/services/searchEngine"
	"app/storage"
	"log"
)

func main() {

	var args = new(cmd.Args)
	args.Parse()

	var config = configLoader.Load(args)
	var storage = storage.GetStorage(config.Storage)
	var searchEngine = &searchEngineService.SearchEngine{Storage: storage}

	if args.Source != "" {
		var provider = provider.StrategyFactoryProvider(args.Source)

		var dataLoader = dataLoaderService.Loader{
			Storage:  storage,
			Provider: provider,
		}
		dataLoader.Run()
	}

	log.Println("Run web server")
	server.Run(config.WebServer, searchEngine)
}
