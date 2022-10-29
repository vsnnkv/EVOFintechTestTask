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
