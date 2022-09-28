package entity_test

import (
	"assert"
	"testing"
)

func TestOrder(t *testing.T) {
	order := entity.Order{
		ID:         "1",
		Price:      100,
		Tax:        10,
		TotalPrice: 110,
	}
	assert.Equal(t, order.ID, "1")
	assert.Equal(t, order.Price, 100.0)
	assert.Equal(t, order.Tax, 10.0)
	assert.Equal(t, order.TotalPrice, 110.0)
}
