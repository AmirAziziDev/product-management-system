package handlers

import (
	"net/http"
	"sync"

	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductsResponse struct {
	Data []models.Product `json:"data"`
	Meta struct {
		Total    int `json:"total"`
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	} `json:"meta"`
}

func ListProducts(logger *zap.Logger, repo repositories.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("ListProducts handler called")

		page := c.GetInt("page")
		pageSize := c.GetInt("page_size")

		// Run both queries concurrently
		var wg sync.WaitGroup
		var total int
		var products []models.Product
		var countErr, productsErr error

		wg.Add(2)

		go func() {
			defer wg.Done()
			total, countErr = repo.GetProductsCount()
		}()

		go func() {
			defer wg.Done()
			products, productsErr = repo.GetProducts(page, pageSize)
		}()

		wg.Wait()

		if countErr != nil {
			logger.Error("Failed to get total products count", zap.Error(countErr))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get products count",
			})
			return
		}

		if productsErr != nil {
			logger.Error("Failed to fetch products from repository", zap.Error(productsErr))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch products",
			})
			return
		}

		response := ProductsResponse{
			Data: products,
		}
		response.Meta.Total = total
		response.Meta.Page = page
		response.Meta.PageSize = pageSize

		c.JSON(http.StatusOK, response)
	}
}
