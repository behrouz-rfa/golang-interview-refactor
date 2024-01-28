//#go:build integration

package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"interview/pkg/core/dto"
)

var itemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

func (s *CartTestSuite) TestEmptyGetCartData() {
	sessionID := uuid.New().String()

	items := s.cartService.GetCartData(sessionID)
	assert.Empty(s.T(), items)

}

func (s *CartTestSuite) TestAddItemToCartData() {
	sessionID := uuid.New().String()

	err := s.cartService.AddItemToCart(&dto.CartItemForm{
		Product:  "shoe",
		Quantity: "100",
	}, sessionID)
	assert.Nil(s.T(), err)
	items := s.cartService.GetCartData(sessionID)
	assert.Len(s.T(), items, 1)
}

func (s *CartTestSuite) TestAddItemInvalidToCartData() {
	sessionID := uuid.New().String()

	err := s.cartService.AddItemToCart(&dto.CartItemForm{
		Product:  "var",
		Quantity: "100",
	}, sessionID)
	assert.NotNil(s.T(), err)
	items := s.cartService.GetCartData(sessionID)
	assert.Len(s.T(), items, 0)
}

func (s *CartTestSuite) TestDeleteCartItem() {
	sessionID := uuid.New().String()

	err := s.cartService.AddItemToCart(&dto.CartItemForm{
		Product:  "shoe",
		Quantity: "100",
	}, sessionID)
	assert.Nil(s.T(), err)

	items := s.cartService.GetCartData(sessionID)
	assert.Len(s.T(), items, 1)

	value, Ok := items[0]["ID"]
	assert.True(s.T(), Ok)

	err = s.cartService.DeleteCartItem(fmt.Sprint(value), sessionID)
	assert.Nil(s.T(), err)

	items = s.cartService.GetCartData(sessionID)
	assert.Len(s.T(), items, 0)
}
