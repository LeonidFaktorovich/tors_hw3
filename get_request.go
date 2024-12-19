package main

import (
	"encoding/json"
	"net/http"
)

type GetResponse struct {
	Value map[string]string `json:"value"`
}

func Get(state *State, w http.ResponseWriter, r *http.Request) {
	state.global_mutex.RLock()
	defer state.global_mutex.RUnlock()

	response := GetResponse{Value: make(map[string]string)}

	for key, value := range state.kv_storage.storage {
		response.Value[key] = value.data
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}
