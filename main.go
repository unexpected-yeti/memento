package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/unexpected-yeti/memento/app"
	"github.com/unexpected-yeti/memento/config"
)

func main() {

	application := app.App{}
	configuration := config.GetConfig()

	application.Initialize(configuration)

	address := ":" + strconv.FormatUint(uint64(configuration.Port), 10)

	log.Print("memento started and listening on localhost", address)

	err := http.ListenAndServe(address, application.Handler)
	if err != nil {
		log.Fatal(err)
	}
}
