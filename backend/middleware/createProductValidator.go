package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateProductRequest Uses default Gin (go-playground) validator only.
type CreateProductRequest struct {
	Code        int     `json:"code"             binding:"required,min=1"`
	Name        string  `json:"name"             binding:"required,min=1"`
	Description *string `json:"description"      binding:"omitempty"`
	ProductType int     `json:"product_type_id"  binding:"required,min=1"`
	ColorIDs    []int   `json:"color_ids"        binding:"required,unique,dive,gt=0"`
}

func ValidateCreateProductRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": gin.H{
					"global":  "invalid request body",
					"details": strings.TrimSpace(err.Error()),
				},
			})
			c.Abort()
			return
		}

		if req.Description != nil {
			trimmed := strings.TrimSpace(*req.Description)
			req.Description = &trimmed
		}

		c.Set("createProductRequest", req)
		c.Next()
	}
}
