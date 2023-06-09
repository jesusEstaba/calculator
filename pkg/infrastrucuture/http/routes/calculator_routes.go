package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/handlers"
)

type CalculatorRoutes struct {
	handler *handlers.CalculatorHandler
}

func NewCalculatorRoutes(handler *handlers.CalculatorHandler) *CalculatorRoutes {
	return &CalculatorRoutes{
		handler,
	}
}

func (r *CalculatorRoutes) RegisterCalculatorRoutes(secret *gin.RouterGroup) {
	secret.POST("/calculate", r.handler.Calculate)
	secret.POST("/records", r.handler.SearchRecords)
}
