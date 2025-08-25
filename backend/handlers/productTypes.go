package handlers

import (
	"net/http"

	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductTypesResponse struct {
	Data []models.ProductType `json:"data"`
}

func ListProductTypes(logger *zap.Logger, repo repositories.ProductTypeRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("ListProductTypes handler called")

		productTypes, err := repo.GetProductTypes()
		if err != nil {
			logger.Error("Failed to fetch product types from repository", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch product types",
			})
			return
		}

		response := ProductTypesResponse{
			Data: productTypes,
		}

		c.JSON(http.StatusOK, response)
	}
}
