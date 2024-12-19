package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	state *State
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	Get(h.state, w, r)
}

func (h *Handler) PatchHandler(w http.ResponseWriter, r *http.Request) {
	Patch(h.state, w, r)
}

func (h *Handler) UpdateSettingsHandler(w http.ResponseWriter, r *http.Request) {
	UpdateSettings(h.state, w, r)
}

func HttpDaemon(state *State, port int) {
	r := mux.NewRouter()

	h := Handler{state: state}

	r.HandleFunc("/read", h.GetHandler).Methods("GET")
	r.HandleFunc("/update", h.PatchHandler).Methods("PATCH")
	r.HandleFunc("/update_settings", h.UpdateSettingsHandler).Methods("POST")

	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("Start http server on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Error in http server: %v", err)
	}
}
