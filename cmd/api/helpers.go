package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]interface{}

func writeJSON(w http.ResponseWriter, status int, data envelope) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
