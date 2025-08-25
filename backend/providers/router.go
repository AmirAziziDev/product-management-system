package providers

import (
	"github.com/AmirAziziDev/product-management-system/routes"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// NewRouter creates a new Gin router with all routes configured
func NewRouter(logger *zap.Logger, db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	routes.SetupRoutes(router, logger, db)
	return router
}
