package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/entities"
	"github.com/jesusEstaba/calculator/pkg/usecase"
	"net/http"
)

type CalculatorHandler struct {
	calculator *usecase.CalculatorUseCase
}

func NewCalculatorHandler(
	calculator *usecase.CalculatorUseCase,
) *CalculatorHandler {
	return &CalculatorHandler{
		calculator,
	}
}

// Calculate
// @Tags Calculator
// @Summary Perform a operation
// @Accept json
// @Produce json
// @Param request body domain.Calculation true "query params"
// @Success 201 {object} domain.CalculationResult
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /calculate [post]
func (h *CalculatorHandler) Calculate(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, entities.ErrorResponse{Error: "you cant not perform this action"})
		return
	}

	var operation domain.Calculation
	if err := ctx.BindJSON(&operation); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	result, err := h.calculator.Calculate(userID, &operation)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
