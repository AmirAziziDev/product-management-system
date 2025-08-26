package handlers

import (
	"net/http"

	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ColorsResponse struct {
	Data []models.Color `json:"data"`
}

func ListColors(logger *zap.Logger, repo repositories.ColorRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("ListColors handler called")

		colors, err := repo.GetColors()
		if err != nil {
			logger.Error("Failed to fetch colors from repository", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch colors",
			})
			return
		}

		response := ColorsResponse{
			Data: colors,
		}

		c.JSON(http.StatusOK, response)
	}
}
