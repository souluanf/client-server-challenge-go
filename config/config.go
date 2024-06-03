package config

import "time"

const (
	ServerAddress = "localhost:8080"
	ClientTimeout = 300 * time.Millisecond
	DataFolder    = "data/"
	DbFilePath    = "quotation.db"
	DbTimeout     = 10 * time.Millisecond
	ApiUrl        = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	ApiTimeout    = 200 * time.Millisecond
	FileName      = "quotation.txt"
)
