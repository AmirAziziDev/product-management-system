package products

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AmirAziziDev/product-management-system/handlers"
	"github.com/AmirAziziDev/product-management-system/providers"
	"github.com/AmirAziziDev/product-management-system/repositories"
	"github.com/AmirAziziDev/product-management-system/tests/integration/shared"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

func TestProductsEndpointHappyPath(t *testing.T) {
	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16",
		postgres.WithDatabase("product_management_test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	require.NoError(t, err)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			t.Logf("failed to terminate container: %s", err)
		}
	}()

	host, err := postgresContainer.Host(ctx)
	require.NoError(t, err)

	port, err := postgresContainer.MappedPort(ctx, "5432")
	require.NoError(t, err)

	dsn := fmt.Sprintf("host=%s port=%s user=postgres password=testpass dbname=product_management_test sslmode=disable",
		host, port.Port())

	db, err := sqlx.Connect("postgres", dsn)
	require.NoError(t, err)
	defer db.Close()

	err = shared.InitializeProductsSchema(db)
	require.NoError(t, err)

	err = shared.SeedProductData(db)
	require.NoError(t, err)

	logger, _ := zap.NewDevelopment()
	productRepo := repositories.NewProductRepository(db)
	productTypeRepo := repositories.NewProductTypeRepository(db)
	colorRepo := repositories.NewColorRepository(db)

	gin.SetMode(gin.TestMode)
	router := providers.NewRouter(logger, productRepo, productTypeRepo, colorRepo)

	req, err := http.NewRequest("GET", "/api/v1/products?page=1&page_size=20", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response handlers.ProductsResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response.Data)
	assert.Equal(t, 1, response.Meta.Page)
	assert.Equal(t, 20, response.Meta.PageSize)
	assert.Equal(t, 10, response.Meta.Total) // We seeded 10 products

	assert.LessOrEqual(t, len(response.Data), 20)
	assert.Greater(t, len(response.Data), 0)

	firstProduct := response.Data[0]
	assert.Greater(t, firstProduct.ID, 0)
	assert.Greater(t, firstProduct.Code, 0)
	assert.NotEmpty(t, firstProduct.Name)
	assert.NotZero(t, firstProduct.CreatedAt)

	if len(response.Data) > 1 {
		for i := 0; i < len(response.Data)-1; i++ {
			assert.True(t, response.Data[i].CreatedAt.After(response.Data[i+1].CreatedAt) ||
				response.Data[i].CreatedAt.Equal(response.Data[i+1].CreatedAt),
				"Products should be ordered by created_at DESC")
		}
	}

	t.Logf("Integration test passed! Retrieved %d products (total: %d)",
		len(response.Data), response.Meta.Total)
}
