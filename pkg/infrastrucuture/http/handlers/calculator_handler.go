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
	records    *usecase.SearchUserRecordsUseCase
}

func NewCalculatorHandler(
	calculator *usecase.CalculatorUseCase,
	records *usecase.SearchUserRecordsUseCase,
) *CalculatorHandler {
	return &CalculatorHandler{
		calculator,
		records,
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

// SearchRecords
// @Tags Records
// @Summary Get paginated records by search term
// @Accept json
// @Produce json
// @Param request body domain.RecordSearch true "query params"
// @Success 201 {object} entities.SearchRecordsResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /records [post]
func (h *CalculatorHandler) SearchRecords(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, entities.ErrorResponse{Error: "you cant not perform this action"})
		return
	}

	var search domain.RecordSearch
	if err := ctx.BindJSON(&search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	search.UserID = userID

	result, err := h.records.Search(search)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entities.SearchRecordsResponse{Records: result})
}
