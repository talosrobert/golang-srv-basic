package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/talosrobert/golang-srv-basic/internal/advertisments"
)

type application struct {
	ads *advertisments.AdvertismentInventory
}

func newApplication() *application {
	return &application{
		ads: advertisments.NewAdvertismentInventory(),
	}
}

func (appl *application) createAdvertismentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Tags []string `json:"tags"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ad, err := appl.ads.CreateAdvertisment(input.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully created an advertisment with ID: %v", ad.ID)
}

func (appl *application) getAdvertismentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid Advertisment ID", http.StatusBadRequest)
		return
	}

	ad, err := appl.ads.GetAdvertismentByID(id)
	if err != nil {
		http.Error(w, "Invalid Advertisment ID", http.StatusNotFound)
		return
	}

	adjs, err := json.Marshal(ad)
	if err != nil {
		http.Error(w, "Invalid Advertisment ID", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(adjs)
}
