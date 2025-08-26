package providers

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDatabaseConfig creates database configuration from environment variables
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnvOrDefault("POSTGRES_HOST", "localhost"),
		Port:     getEnvOrDefault("POSTGRES_PORT", "5432"),
		User:     getEnvOrDefault("POSTGRES_USER", "product_user"),
		Password: getEnvOrDefault("POSTGRES_PASSWORD", "product_password"),
		DBName:   getEnvOrDefault("POSTGRES_DB", "product_management"),
		SSLMode:  getEnvOrDefault("POSTGRES_SSLMODE", "disable"),
	}
}

// NewDatabase creates a new database connection
func NewDatabase(config *DatabaseConfig, logger *zap.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database", zap.Error(err))
		return nil, err
	}

	logger.Info("Successfully connected to database",
		zap.String("host", config.Host),
		zap.String("port", config.Port),
		zap.String("dbname", config.DBName))

	return db, nil
}

// getEnvOrDefault gets environment variable or returns default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
