package routes

import (
	"github.com/AmirAziziDev/product-management-system/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/healthz", handlers.HealthCheck)
}
