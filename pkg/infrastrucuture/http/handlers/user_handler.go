package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/entities"
	"github.com/jesusEstaba/calculator/pkg/usecase"
	"net/http"
)

type UserHandler struct {
	register *usecase.RegisterUserUseCase
}

func NewUserHandler(
	register *usecase.RegisterUserUseCase,
) *UserHandler {
	return &UserHandler{
		register,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var create entities.CreateUser
	if err := ctx.BindJSON(&create); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	created, err := h.register.RegisterUser(create.Username, create.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}
