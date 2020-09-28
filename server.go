package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	r := mux.NewRouter()
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)

	server := http.Server{
		Addr:              ":2009",
		Handler:           r,
	}

	log.Println("starting...")
	log.Fatal(server.ListenAndServe())
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"active"}`))
}
