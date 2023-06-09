package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/entities"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/third_party"
	"net/http"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, entities.ErrorResponse{Error: "Unauthorized"})
			return
		}

		repo := third_party.NewTokenRepositoryImplementation()
		userID, err := repo.Verify(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, entities.ErrorResponse{Error: "Unauthorized"})
			return
		}

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
