package controllers

import (
	"EVOFintechTestTask/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
	service services.TransactionServiceInterface
}

func NewTransactionController(s services.TransactionServiceInterface) *TransactionController {
	return &TransactionController{
		service: s,
	}
}

func (controller *TransactionController) SaveTransactions(c *gin.Context) {
	err := controller.service.SaveData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func (controller *TransactionController) Get(c *gin.Context) {
	receivedDataFromFilters := services.NewReceivedFiltersFilters(c.QueryArray("transaction_id"),
		c.QueryArray("terminal_id"), c.Query("status"), c.Query("payment_type"),
		c.Query("from"), c.Query("to"), c.Query("payment_narrative"))

	transactions, err := controller.service.GetFilteredData(receivedDataFromFilters)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, transactions)
}
