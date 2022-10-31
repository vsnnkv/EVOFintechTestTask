package presentation

import (
	"EVOFintechTestTask/presentation/controllers"
	"EVOFintechTestTask/repository"
	"EVOFintechTestTask/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	transactionController *controllers.TransactionController
}

func New(t *controllers.TransactionController) *Router {
	return &Router{transactionController: t}
}

func (r *Router) CreateRouter() {
	router := gin.Default()

	router.GET("/saveTransactions", r.transactionController.SaveTransactions)
	router.GET("/getSMTH", r.transactionController.Get)

	router.Run(":8080")
}

func InitRoutes() {
	repository := repository.DB{}

	transactionService := services.NewTransactionService(repository)
	transactionController := controllers.NewTransactionController(transactionService)

	router := New(transactionController)
	router.CreateRouter()
}
