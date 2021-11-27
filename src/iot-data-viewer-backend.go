package main

import (
	"log"
	"net/http"

	api "iot-data-viewer-backend/src/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const apiVersion = "/api/v1"

func main() {
	log.Println("starting application")

	router := mux.NewRouter()
	router.HandleFunc(apiVersion+"/health", api.HealthStatus).Methods("GET")
	router.HandleFunc(apiVersion+"/data", api.SaveData).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	server := &http.Server{Addr: ":3001", Handler: handler}

	log.Fatal(server.ListenAndServe())

}
