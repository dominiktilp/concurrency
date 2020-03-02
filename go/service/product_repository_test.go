package service_test

import (
	"concurrency/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepository(t *testing.T) {
	repo, err := service.NewProductRepository("http://0.0.0.0:9999")
	assert.NoError(t, err)

	t.Run("get all products", func(t *testing.T) {
		result, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, result, 1000)
	})

	t.Run("get product by ID", func(t *testing.T) {
		products, err := repo.GetAll()
		assert.NoError(t, err)
		expected := products[10]

		result, err := repo.Get(expected.Id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})

	t.Run("get all recommended products", func(t *testing.T) {
		result, err := repo.GetRecommendedAll()
		assert.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("get recommended product by ID", func(t *testing.T) {
		recommended, err := repo.GetRecommendedAll()
		assert.NoError(t, err)
		expected := recommended[5]

		result, err := repo.GetRecommended(expected.Id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})
}
