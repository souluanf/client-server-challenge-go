package utils

import (
	"client-server-challenge-go/config"
	"fmt"
	"os"
	"time"
)

func SaveToFile(bid string) error {
	if _, err := os.Stat(config.DataFolder); os.IsNotExist(err) {
		err := os.Mkdir(config.DataFolder, 0755)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(config.DataFolder+config.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file: ", err)
		}
	}(file)

	timestamp := time.Now().Format(time.DateTime)
	_, err = fmt.Fprintf(file, "%s - Dólar: %s\n", timestamp, bid)
	return err
}
