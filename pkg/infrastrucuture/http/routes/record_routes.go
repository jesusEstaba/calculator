package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/handlers"
)

type RecordRoutes struct {
	handler *handlers.RecordHandler
}

func NewRecordRoutes(handler *handlers.RecordHandler) *RecordRoutes {
	return &RecordRoutes{
		handler,
	}
}

func (r *RecordRoutes) RegisterRecordRoutes(secret *gin.RouterGroup) {
	secret.POST("/records", r.handler.SearchRecords)
	secret.DELETE("/records/:ID", r.handler.DeleteRecord)
}
