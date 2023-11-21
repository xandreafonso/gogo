package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIdisBlank(t *testing.T) {
	order := Order{} // Id: "1", Price: 1, Tax: 1

	// if order.Validate() == nil {
	// 	t.Error("id is required")
	// }

	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{Id: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{Id: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{Id: "123", Price: 10.0, Tax: 1.0}

	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)

	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
