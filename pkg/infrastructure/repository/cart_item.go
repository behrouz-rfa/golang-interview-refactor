package repository

import (
	"log"

	"interview/pkg/core/entity"
)

func (u Repository) DeleteCartItem(cartItemEntity *entity.CartItem) error {
	return u.db.Delete(cartItemEntity).Error
}

func (u Repository) SaveCartItem(e *entity.CartItem) error {
	err := u.db.Save(e).Error
	log.Println(err)
	return err
}

func (u Repository) CreateCartItem(cartEntity *entity.CartItem) error {
	return u.db.Create(cartEntity).Error
}

func (u Repository) GetCartItemsByCartID(id uint) (cartItems []entity.CartItem, err error) {
	err = u.db.Where("cart_id = ?", id).Find(&cartItems).Error
	return
}

func (u Repository) GetCartItemByID(id uint) (cartItem entity.CartItem, err error) {
	err = u.db.Where("id = ?", id).First(&cartItem).Error
	return
}

func (u Repository) GetCartEntityByProductNameAndID(id uint, productName string) (cartItem entity.CartItem, err error) {
	err = u.db.Where(" cart_id = ? and product_name  = ?", id, productName).First(&cartItem).Error
	return
}
