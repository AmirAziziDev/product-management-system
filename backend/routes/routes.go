package routes

import (
	"github.com/AmirAziziDev/product-management-system/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func SetupRoutes(router *gin.Engine, logger *zap.Logger, db *sqlx.DB) {
	router.GET("/healthz", handlers.HealthCheck)
	router.GET("/api/v1/products", handlers.ListProducts(logger, db))
}
