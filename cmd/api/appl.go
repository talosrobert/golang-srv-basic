package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/talosrobert/golang-srv-basic/internal/advertisments"
)

type application struct {
	cfg    config
	logger *log.Logger
	ads    *advertisments.AdvertismentInventory
}

func newApplication(cfg config) *application {
	return &application{
		cfg:    cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		ads:    advertisments.NewAdvertismentInventory(),
	}
}

func (appl *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthcheck", appl.healthcheckHandler)
	mux.HandleFunc("GET /v1/ad/{id}", appl.getAdvertismentHandler)
	mux.HandleFunc("POST /v1/ad", appl.createAdvertismentHandler)
	return mux
}

func (appl *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": appl.cfg.env,
		"version":     version,
	}

	err := writeJSON(w, http.StatusOK, envelope{"healthcheck": data})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (appl *application) createAdvertismentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Tags []string `json:"tags"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ad, err := appl.ads.CreateAdvertisment(input.Tags)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"ad": ad})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (appl *application) getAdvertismentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ad, err := appl.ads.GetAdvertismentByID(id)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"ad": ad})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}
