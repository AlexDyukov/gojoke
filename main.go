// gojoke is REST api service to store & edit & generate jokes
package main

import (
	"log"
	"net/http"

	"github.com/alexdyukov/gojoke/api"
	"github.com/alexdyukov/gojoke/db"
)

func main() {
	jdb, err := db.Init()
	if err != nil {
		panic(err)
	}
	router := api.ApiRouter(jdb)
	log.Fatal(http.ListenAndServe(":8080", router))
}
