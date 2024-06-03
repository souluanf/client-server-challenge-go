package main

import (
	"client-server-challenge-go/config"
	"client-server-challenge-go/utils"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db, err := utils.InitDatabase()
	if err != nil {
		log.Fatal("Error initializing the database: ", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing the database: ", err)
		}
	}(db)

	log.Println("Starting server on:", config.ServerAddress)
	http.HandleFunc("/cotacao", cotacaoHandler(db))
	if err := http.ListenAndServe(config.ServerAddress, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func cotacaoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), config.ApiTimeout)
		defer cancel()

		data, err := utils.FetchDataFromApi(ctx)
		if err != nil {
			http.Error(w, "Error getting data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		floatValue, err := strconv.ParseFloat(data.Bid, 64)
		if err != nil {
			floatValue = 0.0
		}

		if err := utils.SaveData(db, floatValue); err != nil {
			http.Error(w, "Error saving data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Error encoding data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
