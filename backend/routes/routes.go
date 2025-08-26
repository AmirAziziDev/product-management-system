package routes

import (
	"github.com/AmirAziziDev/product-management-system/handlers"
	"github.com/AmirAziziDev/product-management-system/middleware"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(router *gin.Engine, logger *zap.Logger, productRepo repositories.ProductRepository, productTypeRepo repositories.ProductTypeRepository, colorRepo repositories.ColorRepository) {
	router.GET("/healthz", handlers.HealthCheck())
	router.GET("/api/v1/products", middleware.ValidateProductsQuery(), handlers.ListProducts(logger, productRepo))
	router.GET("/api/v1/product-types", handlers.ListProductTypes(logger, productTypeRepo))
	router.GET("/api/v1/colors", handlers.ListColors(logger, colorRepo))
}
