package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/unexpected-yeti/memento/app"
	"github.com/unexpected-yeti/memento/app/database"
	"github.com/unexpected-yeti/memento/config"
)

func main() {

	application := app.App{}
	configuration := config.GetConfig()
	database := database.NewFSDatastore(configuration.DataDir)

	application.Initialize(configuration, database)

	address := fmt.Sprintf(":%d", configuration.Port)

	LogYeti()
	log.Print("Memento started and listening on localhost", address)
	log.Print("Using FSDatastore, serving data from ", configuration.DataDir)

	err := http.ListenAndServe(address, application.Handler)
	if err != nil {
		log.Fatal(err)
	}
}
