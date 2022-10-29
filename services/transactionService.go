package services

import (
	"EVOFintechTestTask/config"
	"EVOFintechTestTask/repository"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TransactionServiceInterface interface {
	SaveData() error
}

type TransactionService struct {
	db repository.DBInterface
}

func NewTransactionService(db repository.DB) *TransactionService {
	return &TransactionService{db: &db}
}

func (service *TransactionService) SaveData() error {
	err, fileName := downloadData()
	if err != nil {
		return err
	}

	err = service.db.SaveDataFromCSVToDB(fileName)
	return err
}

func downloadData() (error, string) {
	cfg := config.Get()
	url := cfg.DownloadUrl + cfg.FileId + cfg.DownloadUrlEnding
	fileName := "file.csv"

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return err, ""
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	fmt.Println(n, "bytes downloaded")

	return err, fileName
}