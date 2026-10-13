package httpapi

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hoangbinh25/spa-booking/internal/branch"
	"github.com/hoangbinh25/spa-booking/internal/health"
)

// Handlers are the HTTP handlers made available by the application.
type Handlers struct {
	Health *health.Handler
	Branch *branch.Handler
}

// NewRouter registers application routes and middleware.
func NewRouter(handlers Handlers) *gin.Engine {
	router := gin.New()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	router.GET("/health", handlers.Health.Check)

	api := router.Group("/api/v1")
	{
		api.GET("/branches", handlers.Branch.List)
		api.GET("/branches/:id", handlers.Branch.GetByID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
