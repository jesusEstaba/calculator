package main

import (
	"github.com/jesusEstaba/calculator/internal"
	"github.com/jesusEstaba/calculator/internal/database"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/handlers"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/http/routes"
	httprepo "github.com/jesusEstaba/calculator/pkg/infrastrucuture/repositories/http"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/repositories/persistence"
	"github.com/jesusEstaba/calculator/pkg/infrastrucuture/repositories/third_party"
	"github.com/jesusEstaba/calculator/pkg/usecase"
	"github.com/sirupsen/logrus"
)

// @title     Calculator API
// @version   1.0

// @host      localhost:8080
// @BasePath  /api/calculator/v1

func main() {
	api := dependencies()

	healthRoutes := routes.NewHealthCheckRoutes()
	swaggerRoutes := routes.NewSwaggerDocsRoutes()
	userRoutes := routes.NewRoutes(api.userHandler)
	calculatorRoutes := routes.NewCalculatorRoutes(api.calculatorHandler)

	routesGroup := http.RoutesGroup{
		HealthCheckRoutes: healthRoutes,
		SwaggerRoutes:     swaggerRoutes,
		UserRoutes:        userRoutes,
		CalculatorRoutes:  calculatorRoutes,
	}

	r := http.NewRouter(routesGroup)
	logrus.Fatal(r.Run(":8080"))
}

type API struct {
	userHandler       *handlers.UserHandler
	calculatorHandler *handlers.CalculatorHandler
}

func dependencies() *API {
	db, err := database.NewDatabase(&internal.Config)
	if err != nil {
		logrus.Fatal(err)
	}

	userRepoImpl := persistence.NewUserRepository(db)
	passwdRepoImpl := third_party.NewPasswordRepository()
	tokenRepoImpl := third_party.NewTokenRepositoryImplementation()
	randomRepoImpl := httprepo.NewRandomStringRepositoryImplementation()
	operationRepoImpl := persistence.NewOperationRepositoryImplementation(db)

	registerUseCase := usecase.NewRegisterUserUseCase(
		userRepoImpl,
		passwdRepoImpl,
	)

	loginUseCase := usecase.NewLoginUseCase(
		userRepoImpl,
		passwdRepoImpl,
		tokenRepoImpl,
	)

	userHandler := handlers.NewUserHandler(
		registerUseCase,
		loginUseCase,
	)

	calculatorUseCase := usecase.NewCalculatorUseCase(
		userRepoImpl,
		operationRepoImpl,
		randomRepoImpl,
	)

	searchRecords := usecase.NewSearchUserRecordsUseCase(
		operationRepoImpl,
	)

	calculatorHandler := handlers.NewCalculatorHandler(
		calculatorUseCase,
		searchRecords,
	)

	return &API{
		userHandler,
		calculatorHandler,
	}
}
