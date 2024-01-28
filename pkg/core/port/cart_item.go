package port

import (
	"interview/pkg/core/dto"
	"interview/pkg/core/entity"
)

type CartService interface {
	AddItemToCart(*dto.CartItemForm, string) error
	DeleteCartItem(string, string) error
	GetCartData(sessionID string) (items []map[string]interface{})
}

type CartItemRepository interface {
	DeleteCartItem(cartItem *entity.CartItem) error
	GetCartItemByID(id uint) (entity.CartItem, error)
	SaveCartItem(cartItem *entity.CartItem) error
	GetCartItemsByCartID(cartID uint) ([]entity.CartItem, error)
	GetCartEntityByProductNameAndID(id uint, productName string) (entity.CartItem, error)
	CreateCartItem(cartItem *entity.CartItem) error
}

type CartEntityRepository interface {
	GetCartEntityByStatusAndSessionID(status, sessionID string) (entity.Cart, error)
	CreateCartEntity(cartEntity *entity.Cart) error
}
