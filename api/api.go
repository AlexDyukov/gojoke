package api

import (
	"net/http"

	"github.com/alexdyukov/gojoke/db"
	"github.com/gorilla/mux"
)

func ApiRouter(jdb db.JokeDatabase) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		generateJoke(jdb, w, r)
	}).Methods("GET")
	router.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		listJokes(jdb, w, r)
	}).Methods("GET")
	router.HandleFunc("/joke/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		getJoke(jdb, w, r)
	}).Methods("GET")
	router.HandleFunc("/joke/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		updateJoke(jdb, w, r)
	}).Methods("POST")
	router.HandleFunc("/joke/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		deleteJoke(jdb, w, r)
	}).Methods("DELETE")
	return router
}

func generateJoke(jdb db.JokeDatabase, w http.ResponseWriter, r *http.Request) {
}
func listJokes(jdb db.JokeDatabase, w http.ResponseWriter, r *http.Request) {
}
func getJoke(jdb db.JokeDatabase, w http.ResponseWriter, r *http.Request) {
}
func updateJoke(jdb db.JokeDatabase, w http.ResponseWriter, r *http.Request) {
}
func deleteJoke(jdb db.JokeDatabase, w http.ResponseWriter, r *http.Request) {
}
