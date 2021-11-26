package api

import (
	"net/http"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
