package providers

import (
	"context"
	"errors"
	"net/http"

	"github.com/AmirAziziDev/product-management-system/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewLogger creates a new structured logger
func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

// NewRouter creates a new Gin router with all routes configured
func NewRouter() *gin.Engine {
	router := gin.Default()
	routes.SetupRoutes(router)
	return router
}

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return srv
}

// Run handles the lifecycle management for the HTTP server
func Run(lc fx.Lifecycle, server *http.Server, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting server on :8080")
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Fatal("Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down server...")
			if err := server.Shutdown(ctx); err != nil {
				logger.Error("Server forced to shutdown", zap.Error(err))
				return err
			}
			logger.Info("Server exited gracefully")
			return nil
		},
	})
}
