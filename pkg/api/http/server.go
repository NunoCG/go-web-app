package http

import (
	"first/pkg/api"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	router := mux.NewRouter()
	router.HandleFunc("/clients", api.GetClientsEndPoint).Methods("GET")
	router.HandleFunc("/clients", api.CreateClientEndPoint).Methods("POST")
	http.ListenAndServe(":8000", router)
}
