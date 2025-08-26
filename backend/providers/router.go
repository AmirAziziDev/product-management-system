package providers

import (
	"github.com/AmirAziziDev/product-management-system/middleware"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/AmirAziziDev/product-management-system/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewRouter creates a new Gin router with all routes configured
func NewRouter(logger *zap.Logger, productRepo repositories.ProductRepository, productTypeRepo repositories.ProductTypeRepository, colorRepo repositories.ColorRepository) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	routes.SetupRoutes(router, logger, productRepo, productTypeRepo, colorRepo)
	return router
}
