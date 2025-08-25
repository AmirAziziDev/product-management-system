package handlers

import (
	"net/http"

	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func ListProducts(logger *zap.Logger, db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("ListProducts handler called")

		var products []models.Product
		query := "SELECT id, code, name, description, created_at FROM products ORDER BY id"

		err := db.Select(&products, query)
		if err != nil {
			logger.Error("Failed to fetch products from database", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch products",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": products,
			"meta": gin.H{
				"total":     len(products),
				"page":      1,
				"page_size": len(products),
			},
		})
	}
}
