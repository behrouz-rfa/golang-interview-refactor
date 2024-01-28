//go:build integration

package sql

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"interview/pkg/core/entity"
)

var itemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

func (s *SqlTestSuite) TestCreateCartCreate() {
	sessionID := uuid.New().String()
	cartEntity := entity.Cart{
		Total:     2,
		SessionID: sessionID,
		Status:    entity.CartOpen,
	}

	err := s.cartEntityRepo.CreateCartEntity(&cartEntity)
	assert.Nil(s.T(), err)

	err = s.cartItemRepo.CreateCartItem(&entity.CartItem{
		CartID:      cartEntity.ID,
		ProductName: "shoe",
		Quantity:    10,
		Price:       10 * 100,
	})
	assert.Nil(s.T(), err)

	cartEntityGet, err := s.cartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartOpen, sessionID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), sessionID, cartEntityGet.SessionID)

}

func (s *SqlTestSuite) TestCartItemRepository() {
	// Create a test suite or setup the necessary dependencies

	// Create a new cart item
	cartItem := &entity.CartItem{
		CartID:      1,
		ProductName: "shoe",
		Quantity:    10,
		Price:       10 * 100,
	}

	err := s.cartItemRepo.CreateCartItem(cartItem)
	assert.Nil(s.T(), err)

	// Retrieve the cart item by ID
	retrievedCartItem, err := s.cartItemRepo.GetCartItemByID(cartItem.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), cartItem.CartID, retrievedCartItem.CartID)

	// Retrieve all cart items by cart ID
	cartID := cartItem.CartID
	cartItems, err := s.cartItemRepo.GetCartItemsByCartID(cartID)
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), cartItems)

	// Add more test cases as needed
}

func (s *SqlTestSuite) TestCartRepository() {
	// Create a test suite or setup the necessary dependencies

	// Create a new cart entity
	sessionID := uuid.New().String()
	cartEntity := entity.Cart{
		Total:     2,
		SessionID: sessionID,
		Status:    entity.CartOpen,
	}

	err := s.cartEntityRepo.CreateCartEntity(&cartEntity)
	assert.Nil(s.T(), err)

	// Create a new cart item and associate it with the cart entity
	err = s.cartItemRepo.CreateCartItem(&entity.CartItem{
		CartID:      cartEntity.ID,
		ProductName: "shoe",
		Quantity:    10,
		Price:       10 * 100,
	})
	assert.Nil(s.T(), err)

	// Retrieve the cart entity by status and session ID
	cartEntityGet, err := s.cartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartOpen, sessionID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), sessionID, cartEntityGet.SessionID)

	// Test failure scenarios
	// Attempt to create a cart entity with an existing  ID
	err = s.cartEntityRepo.CreateCartEntity(&cartEntity)
	assert.NotNil(s.T(), err)

	// Attempt to retrieve a non-existent cart entity
	nonExistentSessionID := "non_existent_session_id"
	_, err = s.cartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartOpen, nonExistentSessionID)
	assert.NotNil(s.T(), err)

	// Add more failure test cases as needed
	// Retrieve the cart entity by status and session ID
	_, err = s.cartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartClosed, sessionID)
	assert.NotNil(s.T(), err)

}

func (s *SqlTestSuite) TestSaveCartItem() {
	// Create a test suite or setup the necessary dependencies
	sessionID := uuid.New().String()

	cartEntity := entity.Cart{
		Total:     2,
		SessionID: sessionID,
		Status:    entity.CartOpen,
	}
	err := s.cartEntityRepo.CreateCartEntity(&cartEntity)
	assert.Nil(s.T(), err)

	// Create a new cart item
	cartItem := &entity.CartItem{
		CartID:      cartEntity.ID,
		ProductName: "shoe",
		Quantity:    10,
		Price:       10 * 100,
	}

	// Save the cart item
	err = s.cartItemRepo.SaveCartItem(cartItem)
	assert.Nil(s.T(), err)

	// Retrieve the cart item by ID
	retrievedCartItem, err := s.cartItemRepo.GetCartItemByID(cartItem.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), cartItem.ID, retrievedCartItem.ID)

	// Modify the cart item
	cartItem.Quantity = 5
	cartItem.Price = 5 * 100

	// Save the modified cart item
	err = s.cartItemRepo.SaveCartItem(cartItem)
	assert.Nil(s.T(), err)

	// Retrieve the cart item again
	retrievedCartItem, err = s.cartItemRepo.GetCartItemByID(cartItem.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), cartItem.ID, retrievedCartItem.ID)

	// Test failure scenarios
	// Attempt to Get a cart item with an invalid ID
	retrievedCartItem, err = s.cartItemRepo.GetCartItemByID(999)
	assert.NotNil(s.T(), err)

	// Add more test cases as needed
}
