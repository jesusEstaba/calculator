package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/entities"
	"github.com/jesusEstaba/calculator/pkg/usecase"
	"net/http"
)

type RecordHandler struct {
	records *usecase.SearchUserRecordsUseCase
	delete  *usecase.DeleteRecordUseCase
}

func NewRecordHandler(
	records *usecase.SearchUserRecordsUseCase,
	delete *usecase.DeleteRecordUseCase,
) *RecordHandler {
	return &RecordHandler{
		records,
		delete,
	}
}

// SearchRecords
// @Tags Records
// @Summary Get paginated records by search term
// @Accept json
// @Produce json
// @Param request body domain.RecordSearch true "query params"
// @Param Authorization header string true "Token" required:true
// @Success 200 {object} entities.SearchRecordsResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /records [post]
func (h *RecordHandler) SearchRecords(ctx *gin.Context) {
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

// DeleteRecord
// @Tags Records
// @Summary Delete record
// @Accept json
// @Produce json
// @Param request body domain.RecordSearch true "query params"
// @Param Authorization header string true "Token" required:true
// @Success 204 {object} entities.SearchRecordsResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /records/:ID [delete]
func (h *RecordHandler) DeleteRecord(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, entities.ErrorResponse{Error: "you cant not perform this action"})
		return
	}

	id := ctx.Param("ID")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, entities.ErrorResponse{Error: "empty id"})
		return
	}

	err := h.delete.Delete(userID, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
