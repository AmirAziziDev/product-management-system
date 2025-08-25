package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductsQueryParams defines the validation structure for products list endpoint
type ProductsQueryParams struct {
	Page     int `form:"page" binding:"omitempty,min=1"`
	PageSize int `form:"page_size" binding:"omitempty,min=1,max=100"`
}

// ValidateProductsQuery validates query parameters for products list endpoint
func ValidateProductsQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params ProductsQueryParams

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid query parameters",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		if pageParam := c.Query("page"); pageParam == "0" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid query parameters",
				"details": "Key: 'ProductsQueryParams.Page' Error:Field validation for 'Page' failed on the 'min' tag",
			})
			c.Abort()
			return
		}

		if pageSizeParam := c.Query("page_size"); pageSizeParam == "0" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid query parameters",
				"details": "Key: 'ProductsQueryParams.PageSize' Error:Field validation for 'PageSize' failed on the 'min' tag",
			})
			c.Abort()
			return
		}

		// Apply defaults only if values are zero (not provided)
		if params.Page == 0 {
			params.Page = 1
		}
		if params.PageSize == 0 {
			params.PageSize = 20
		}

		// Set validated parameters in context for handler to use
		c.Set("page", params.Page)
		c.Set("page_size", params.PageSize)

		c.Next()
	}
}
