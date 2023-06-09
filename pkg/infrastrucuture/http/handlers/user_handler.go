package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/entities"
	"github.com/jesusEstaba/calculator/pkg/usecase"
	"net/http"
)

type UserHandler struct {
	register *usecase.RegisterUserUseCase
	login    *usecase.LoginUseCase
}

func NewUserHandler(
	register *usecase.RegisterUserUseCase,
	login *usecase.LoginUseCase,
) *UserHandler {
	return &UserHandler{
		register,
		login,
	}
}

// Register
// @Tags User
// @Summary Create a new user
// @Accept json
// @Produce json
// @Param request body entities.CreateUser true "query params"
// @Success 201 {object} domain.User
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /register [post]
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

// Login
// @Tags User
// @Summary User login
// @Accept json
// @Produce json
// @Param request body entities.LoginUser true "query params"
// @Success 200 {object} entities.LoginResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var create entities.LoginUser
	if err := ctx.BindJSON(&create); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.login.Login(create.Username, create.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entities.LoginResponse{Token: *token})
}
