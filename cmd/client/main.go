package main

import (
	"client-server-challenge-go/config"
	"client-server-challenge-go/utils"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), config.ClientTimeout)
	defer cancel()

	quotation, err := utils.FetchQuotation(ctx, "http://"+config.ServerAddress)
	if err != nil {
		log.Fatal("Error fetching quotation: ", err)
	}

	if err := utils.SaveToFile(quotation); err != nil {
		log.Fatal("Error saving quotation: ", err)
	}
}
