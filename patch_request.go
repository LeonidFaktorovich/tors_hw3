package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PatchBody struct {
	Update map[string]string `json:"update"`
}

type PatchResponse struct {
	Value map[string]string `json:"value"`
}

func Patch(state *State, w http.ResponseWriter, r *http.Request) {
	var body PatchBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error with decode: %v", err.Error()), http.StatusBadRequest)
		return
	}

	fmt.Printf("Start process patch request\n")

	response := PatchResponse{Value: make(map[string]string)}

	state.global_mutex.Lock()
	defer state.global_mutex.Unlock()

	err = WriteWithoutLock(state, body.Update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
