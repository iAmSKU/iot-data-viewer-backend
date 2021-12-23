package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	internal "iot-data-viewer-backend/src/internal"
	model "iot-data-viewer-backend/src/model"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling SaveData")
	decoder := json.NewDecoder(r.Body)
	var connectionConfig model.ConnectionConfig
	err := decoder.Decode(&connectionConfig)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Println(connectionConfig)
		internal.AddModifyConnection(connectionConfig)
		fmt.Fprintf(w, "SaveData - OK\n")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling HealthStatus")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HealthStatus - OK\n")
}
