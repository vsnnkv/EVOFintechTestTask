package services

import (
	"EVOFintechTestTask/models"
	"EVOFintechTestTask/repository"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const urlForDownload = "https://drive.google.com/u/0/uc?id=1IwZ3uUCHGpSL2OoQu4mtbw7Ew3ZamcGB&export=download"

type TransactionServiceInterface interface {
	SaveData() error
	GetFilteredData(receivedData *ReceivedFilters) ([]models.Transaction, error)
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

func (service *TransactionService) GetFilteredData(receivedData *ReceivedFilters) ([]models.Transaction, error) {

	var mapForDB = map[string]interface{}{}
	transactionIds, err := checkId(receivedData.TransactionId)
	if err != nil {
		return nil, err
	} else if len(transactionIds) != 0 {
		mapForDB[transactionField] = transactionIds
	}
	terminalIds, err := checkId(receivedData.TerminalId)
	if err != nil {
		return nil, err
	} else if len(terminalIds) != 0 {
		mapForDB[terminalField] = terminalIds
	}
	err = checkStatus(receivedData.Status)
	if err != nil {
		return nil, err
	} else if len(receivedData.Status) != 0 {
		mapForDB[statusField] = receivedData.Status
	}
	err = checkPaymentType(receivedData.PaymentType)
	if err != nil {
		return nil, err
	} else if len(receivedData.PaymentType) != 0 {
		mapForDB[paymentTypeField] = receivedData.PaymentType
	}

	transactions, err := service.db.FindDataInDB(mapForDB)
	if err != nil {
		return nil, err
	}

	if receivedData.PaymentNarrative != "" {
		transactions = lookForNarrative(transactions, receivedData.PaymentNarrative)
	}

	if receivedData.DatePostFrom != "" && receivedData.DatePostTo != "" {
		timeFrom, _ := getTime(receivedData.DatePostFrom)
		timeTo, _ := getTime(receivedData.DatePostTo)

		transactions = lookForDate(transactions, timeFrom, timeTo)
	}

	return *transactions, nil

}

func checkId(s []string) ([]int, error) {
	var passedId []int

	if len(s) == 0 {
		return passedId, nil
	} else {

		for _, id := range s {
			intId, err := strconv.Atoi(id)
			if err != nil {
				return nil, errors.New("couldn't convert ids")
			}

			passedId = append(passedId, intId)
		}
		return passedId, nil

	}
}

func checkStatus(status string) error {
	switch status {
	case "declined":
		return nil
	case "accepted":
		return nil
	case "":
		return nil
	default:
		return errors.New("invalid status")
	}
}

func checkPaymentType(paymentType string) error {
	switch paymentType {
	case "cash":
		return nil
	case "card":
		return nil
	case "":
		return nil
	default:
		return errors.New("invalid paymentType")
	}
}

func lookForNarrative(transactions *[]models.Transaction, narrative string) *[]models.Transaction {
	var transactionAfterNar []models.Transaction

	for _, t := range *transactions {
		if !strings.Contains(t.PaymentNarrative, narrative) {
			transactionAfterNar = append(transactionAfterNar, t)
		}
	}
	return &transactionAfterNar
}

func lookForDate(transactions *[]models.Transaction, timeFrom int64, timeTo int64) *[]models.Transaction {
	var transactionAfterDate []models.Transaction

	for _, t := range *transactions {
		timeFromDb, err := getTime(t.DatePost)
		if err == nil {
			if timeFromDb >= timeFrom && timeFromDb <= timeTo {
				transactionAfterDate = append(transactionAfterDate, t)
			}
		}
		return nil
	}
	return &transactionAfterDate
}

func downloadData() (error, string) {
	url := urlForDownload
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

func getTime(t string) (int64, error) {
	t = strings.Replace(t, " ", "T", -1)
	t = t + "Z"

	newT, err := time.Parse(time.RFC3339, t)
	unixT := newT.Unix()
	return unixT, err
}
