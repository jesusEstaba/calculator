package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRoutes struct{}

func NewSwaggerDocsRoutes() *SwaggerRoutes {
	return &SwaggerRoutes{}
}

func (r *SwaggerRoutes) RegisterSwaggerRoutes(group *gin.RouterGroup) {
	group.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.InstanceName(docs.SwaggerInfo.InstanceName()),
		),
	)
}
