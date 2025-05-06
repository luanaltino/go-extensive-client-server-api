package handler

import (
	"encoding/json"
	service "go-extensive-client-server-api/server/service"
	"net/http"
)

func FindExchange(w http.ResponseWriter, r *http.Request) {
	result, err := service.FindExchange()
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)

}
