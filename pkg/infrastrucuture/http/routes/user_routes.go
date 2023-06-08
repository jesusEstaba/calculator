package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/handlers"
)

type UserRoutes struct {
	handler *handlers.UserHandler
}

func NewRoutes(handler *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler,
	}
}

func (r *UserRoutes) RegisterUserRoutes(public *gin.RouterGroup) {
	public.POST("/register", r.handler.Register)
}