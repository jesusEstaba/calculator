package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/routes"
)

type Router interface {
	Run(addr ...string) error
}

func NewRouter(routes RoutesGroup) Router {
	route := gin.Default()

	root := route.Group("/")
	routes.HealthCheckRoutes.RegisterHealthRoutes(root)
	routes.SwaggerRoutes.RegisterSwaggerRoutes(root)

	service := route.Group("api/calculator/v1")
	routes.UserRoutes.RegisterUserRoutes(service)

	return route
}

type RoutesGroup struct {
	HealthCheckRoutes *routes.HealthRoutes
	SwaggerRoutes     *routes.SwaggerRoutes
	UserRoutes        *routes.UserRoutes
}
