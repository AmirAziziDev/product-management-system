package handlers

import (
	"errors"
	"net/http"

	"github.com/AmirAziziDev/product-management-system/middleware"
	"github.com/AmirAziziDev/product-management-system/models"
	repoif "github.com/AmirAziziDev/product-management-system/repositories/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

func CreateProduct(logger *zap.Logger, repo repoif.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw, exists := c.Get("createProductRequest")
		if !exists {
			logger.Error("createProductRequest missing from context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		req := raw.(middleware.CreateProductRequest)

		product := models.Product{
			Code:        req.Code,
			Name:        req.Name,
			Description: req.Description,
			ProductType: models.ProductType{ID: req.ProductType},
		}

		_, err := repo.CreateProduct(c.Request.Context(), product, req.ColorIDs)
		if handled := handleCreateProductError(c, logger, err); handled {
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "successfully created product"})
	}
}

func handleCreateProductError(c *gin.Context, logger *zap.Logger, err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, repoif.ErrProductTypeNotFound) {
		writeFieldError(c, http.StatusBadRequest, "product_type_id", "product type does not exist")
		return true
	}
	if errors.Is(err, repoif.ErrColorsNotFound) {
		writeFieldError(c, http.StatusBadRequest, "color_ids", "colors do not exist")
		return true
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
		switch pqErr.Constraint {
		case "products_code_unique", "unique_products_code", "products_code_key":
			writeFieldError(c, http.StatusConflict, "products_code", "products code already exists")
			return true
		case "products_name_unique", "products_name_unique_ci", "products_name_key":
			writeFieldError(c, http.StatusConflict, "products_name", "products name already exists")
			return true
		default:
			c.JSON(http.StatusConflict, gin.H{
				"error": "conflict",
				"data":  gin.H{"unique": "duplicate value"},
			})
			return true
		}
	}

	logger.Error("failed to create product", zap.Error(err))
	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
	return true
}

func writeFieldError(c *gin.Context, status int, field, message string) {
	c.JSON(status, gin.H{
		"errors": gin.H{
			field: message,
		},
	})
}
