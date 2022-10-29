package repository

import (
	"errors"
	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const dsn = "host=localhost user=postgres password=postgres dbname=transactions_database port=5432"

type DBInterface interface {
	SaveDataFromCSVToDB(fileName string) error
}

type DB struct {
}

func (*DB) SaveDataFromCSVToDB(fileName string) error {

	db := connectToDb()

	transactions := readFromFile(fileName)

	err := db.AutoMigrate(&Transaction{})
	if err != nil {
		return err
	}

	result := db.Create(transactions)
	if result.Error != nil {
		return errors.New("data with same id`s already exists")
	}

	return nil
}

func connectToDb() *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db

}

func readFromFile(fileName string) *[]Transaction {
	file, err := os.Open(fileName)

	if err != nil {
		log.Printf(err.Error())
	}

	defer file.Close()

	var transactions []Transaction
	err = gocsv.Unmarshal(file, &transactions)
	if err != nil {
		log.Printf(err.Error())
	}
	return &transactions
}
