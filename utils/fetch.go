package utils

import (
	"client-server-challenge-go/config"
	"client-server-challenge-go/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchDataFromApi(ctx context.Context) (*models.Quotation, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.ApiUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var response map[string]models.Quotation
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print(err)
		return nil, err
	}

	data, ok := response["USDBRL"]
	if !ok {
		return nil, fmt.Errorf("exchange rate not found")
	}
	log.Println(data.Bid)

	return &data, nil
}

func FetchQuotation(ctx context.Context, serverAddress string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/cotacao", serverAddress), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response models.Quotation
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print(err)
		return "", err
	}

	return response.Bid, nil
}
