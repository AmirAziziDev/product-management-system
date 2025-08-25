package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/AmirAziziDev/product-management-system/models"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListProducts(logger *zap.Logger, repo repositories.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("ListProducts handler called")

		page := 1
		pageSize := 20

		if pageParam := c.Query("page"); pageParam != "" {
			if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
				page = p
			}
		}

		if pageSizeParam := c.Query("page_size"); pageSizeParam != "" {
			if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 {
				pageSize = ps
			}
		}

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

		c.JSON(http.StatusOK, gin.H{
			"data": products,
			"meta": gin.H{
				"total":     total,
				"page":      page,
				"page_size": pageSize,
			},
		})
	}
}
