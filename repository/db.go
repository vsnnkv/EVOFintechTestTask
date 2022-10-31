package repository

import (
	"EVOFintechTestTask/models"
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
	FindDataInDB(received map[string]interface{}) (*[]models.Transaction, error)
}

type DB struct {
}

func (*DB) SaveDataFromCSVToDB(fileName string) error {
	db := connectToDb()

	transactions := readFromFile(fileName)

	if db.Migrator().HasTable(&models.Transaction{}) {
		return errors.New("table already created")
	}
	err := db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return err
	}

	result := db.Create(transactions)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*DB) FindDataInDB(received map[string]interface{}) (*[]models.Transaction, error) {
	db := connectToDb()

	var transactions []models.Transaction

	result := db.Where(received).Find(&transactions)

	return &transactions, result.Error
}

func connectToDb() *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db

}

func readFromFile(fileName string) *[]models.Transaction {
	file, err := os.Open(fileName)

	if err != nil {
		log.Printf(err.Error())
	}

	defer file.Close()

	var transactions []models.Transaction
	err = gocsv.Unmarshal(file, &transactions)
	if err != nil {
		log.Printf(err.Error())
	}
	return &transactions
}
