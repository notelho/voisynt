package http

import (
	log "log"
	http "net/http"

	mux "github.com/gorilla/mux"
)

func Serve() {

	router := mux.NewRouter()

	router.HandleFunc("/voice", test).Methods("GET")
	router.HandleFunc("/voice", test).Methods("PUT")
	router.HandleFunc("/voice/{id}", test).Methods("POST")
	router.HandleFunc("/voice/{id}", test).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/views/")))

}

func test(w http.ResponseWriter, r *http.Request) {

}
