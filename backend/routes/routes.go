package routes

import (
	"github.com/AmirAziziDev/product-management-system/handlers"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(router *gin.Engine, logger *zap.Logger, productRepo repositories.ProductRepository) {
	router.GET("/healthz", handlers.HealthCheck)
	router.GET("/api/v1/products", handlers.ListProducts(logger, productRepo))
}
