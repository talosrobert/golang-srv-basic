package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/talosrobert/golang-srv-basic/internal/data"
)

type application struct {
	cfg    config
	logger *log.Logger
	models data.Models
}

func newApplication(cfg config, db *sql.DB) *application {
	return &application{
		cfg:    cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		models: data.NewModels(db),
	}
}

func (appl *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthcheck", appl.healthcheckHandler)
	mux.HandleFunc("GET /v1/item/{id}", appl.getAuctionItemHandler)
	mux.HandleFunc("GET /v1/user/{id}", appl.getAuctionUserHandler)
	mux.HandleFunc("PUT /v1/item/{id}", appl.updateAuctionItemHandler)
	mux.HandleFunc("PUT /v1/user/{id}", appl.updateAuctionUserHandler)
	mux.HandleFunc("POST /v1/item", appl.createAuctionItemHandler)
	mux.HandleFunc("POST /v1/user", appl.createAuctionUserHandler)
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
		StartingPrice float64   `json:"starting_price"`
		ReservePrice  float64   `json:"reserve_price"`
		UserID        uuid.UUID `json:"user_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ai := &data.AuctionItem{
		StartingPrice: input.StartingPrice,
		ReservePrice:  input.ReservePrice,
		Seller:        input.UserID,
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
	}
}

func (appl *application) updateAuctionItemHandler(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		StartingPrice *float64 `json:"starting_price,omitempty"`
		ReservePrice  *float64 `json:"reserve_price,omitempty"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.StartingPrice != nil {
		ai.StartingPrice = *input.StartingPrice
	}

	if input.ReservePrice != nil {
		ai.ReservePrice = *input.ReservePrice
	}

	err = appl.models.AuctionItems.Update(ai)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"auction_item": ai})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (appl *application) createAuctionUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		DisplayName string `json:"display_name"`
		EMail       string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	au := &data.AuctionUser{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DisplayName: input.DisplayName,
		EMail:       input.EMail,
	}

	err = appl.models.AuctionUser.Create(au)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"auction_user": au})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (appl *application) getAuctionUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	au, err := appl.models.AuctionUser.Read(id)
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"auction_user": au})
	if err != nil {
		appl.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (appl *application) updateAuctionUserHandler(w http.ResponseWriter, r *http.Request) {}
