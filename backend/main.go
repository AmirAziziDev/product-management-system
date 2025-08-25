package main

import (
	"github.com/AmirAziziDev/product-management-system/providers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			providers.NewLogger,
			providers.NewDatabaseConfig,
			providers.NewDatabase,
			providers.NewRouter,
			providers.NewHTTPServer,
		),
		fx.Invoke(providers.Run),
		fx.NopLogger, // Disable fx's own logging to avoid conflicts with zap
	).Run()
}
