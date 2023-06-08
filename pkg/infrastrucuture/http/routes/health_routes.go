package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthRoutes struct{}

func NewHealthCheckRoutes() *HealthRoutes {
	return &HealthRoutes{}
}

func (r *HealthRoutes) RegisterHealthRoutes(group *gin.RouterGroup) {
	group.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{"message": "pong"})
	})
}
