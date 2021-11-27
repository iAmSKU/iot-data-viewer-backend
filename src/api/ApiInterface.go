package api

import (
	"net/http"
	"fmt"
	"log"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling SaveData")
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Fprintf(w, "SaveData - OK\n")
}

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling HealthStatus")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HealthStatus - OK\n")
}
