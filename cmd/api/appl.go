package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/talosrobert/golang-srv-basic/internal/auctions"
)

type application struct {
	cfg    config
	logger *log.Logger
	models auctions.Models
}

func newApplication(cfg config) *application {
	return &application{
		cfg:    cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

func (appl *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthcheck", appl.healthcheckHandler)
	mux.HandleFunc("GET /v1/ad/{id}", appl.createAuctionItemHandler)
	mux.HandleFunc("POST /v1/ad", appl.getAuctionItemHandler)
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

func (appl *application) createAuctionItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		startingPrice float64
		reservePrice  float64
		userID        uuid.UUID
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ai := &auctions.AuctionItem{
		StartingPrice: input.startingPrice,
		ReservePrice:  input.reservePrice,
		Seller:        input.userID,
	}

	err = appl.models.AuctionItems.Create(ai)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"auction_item": ai})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (appl *application) getAuctionItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ai, err := appl.models.AuctionItems.Read(id)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"auction_item": ai})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}
