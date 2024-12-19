package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UpdateSettingsBody struct {
	Sleep        uint `json:"sleep"`
	BlockReceive bool `json:"block_receive"`
	BlockSend    bool `json:"block_send"`
}

func UpdateSettings(state *State, w http.ResponseWriter, r *http.Request) {
	var body UpdateSettingsBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error with decode: %v", err.Error()), http.StatusBadRequest)
		return
	}

	fmt.Printf("Start process update settings request\n")

	state.global_mutex.Lock()
	defer state.global_mutex.Unlock()

	state.block_receive = body.BlockReceive
	state.sleep = time.Duration(body.Sleep * 1000 * 1000 * 1000)
	if state.block_send != body.BlockSend {
		state.replication_service.ChangeMode()
	}
	state.block_send = body.BlockSend

	w.Write([]byte{})
}
