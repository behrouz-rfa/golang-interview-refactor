package service

import (
	"errors"
	"gorm.io/gorm"
	"interview/pkg/core/dto"
	"interview/pkg/core/entity"
	"interview/pkg/core/port"
	errHandler "interview/pkg/interface/error"
	"interview/pkg/logger"
	"strconv"
)

var itemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

type cartService struct {
	CartServiceParam
	lg *logger.Entry
}

type CartServiceParam struct {
	CartItemRepo   port.CartItemRepository
	CartEntityRepo port.CartEntityRepository
	lg             *logger.Entry
}

func NewCartService(param CartServiceParam) *cartService {
	return &cartService{CartServiceParam: param, lg: logger.General.Component("CartService")}
}

func (cs cartService) AddItemToCart(addItemForm *dto.CartItemForm, sessionID string) error {
	var isCartNew bool
	cartEntity, err := cs.CartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartOpen, sessionID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			cs.lg.WithError(err).Error("GetCartEntityByStatus not working properly")
			return errHandler.ErrRedirect
		}
		isCartNew = true
		cartEntity = entity.Cart{
			SessionID: sessionID,
			Status:    entity.CartOpen,
		}

		err := cs.CartEntityRepo.CreateCartEntity(&cartEntity)
		if err != nil {
			return errHandler.ErrRedirect
		}
	}

	item, ok := itemPriceMapping[addItemForm.Product]
	if !ok {
		return errHandler.ErrRedirect.Msg("invalid item name")
	}

	quantity, err := strconv.ParseInt(addItemForm.Quantity, 10, 0)
	if err != nil {
		cs.lg.WithError(err).Error("ParseInt addItemForm.Quantity failed")
		return errHandler.ErrRedirect.Msg("invalid quantity")
	}

	var cartItemEntity entity.CartItem
	if isCartNew {
		cartItemEntity = entity.CartItem{
			CartID:      cartEntity.ID,
			ProductName: addItemForm.Product,
			Quantity:    int(quantity),
			Price:       item * float64(quantity),
		}

		err := cs.CartItemRepo.CreateCartItem(&cartItemEntity)
		if err != nil {
			cs.lg.WithError(err).Error("CreateCartItem  failed")
			return errHandler.ErrInternalServerError.Msg("internal error")

		}
		return nil
	}

	cartItemEntity, err = cs.CartItemRepo.GetCartEntityByProductNameAndID(cartEntity.ID, addItemForm.Product)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			cs.lg.WithError(err).Error("GetCartEntityByProductName  failed")

			return errHandler.ErrRedirect.Msg("something wrong")
		}

		cartItemEntity = entity.CartItem{
			CartID:      cartEntity.ID,
			ProductName: addItemForm.Product,
			Quantity:    int(quantity),
			Price:       item * float64(quantity),
		}
		err = cs.CartItemRepo.CreateCartItem(&cartItemEntity)
		if err != nil {
			cs.lg.WithError(err).Error("CreateCartItem  failed")

			return errHandler.ErrInternalServerError.Msg("internal error")

		}
		return nil
	}

	cartItemEntity.Quantity += int(quantity)
	cartItemEntity.Price += item * float64(quantity)
	cs.CartItemRepo.SaveCartItem(&cartItemEntity)

	return nil
}

func (cs cartService) DeleteCartItem(cartItemIDString, sessionID string) error {
	_, err := cs.CartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartOpen, sessionID)
	if err != nil {
		return err
	}
	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		cs.lg.WithError(err).Error("DeleteCartItem strconv failed")
		return errHandler.ErrRedirect
	}

	item, err := cs.CartItemRepo.GetCartItemByID(uint(cartItemID))
	if err != nil {
		cs.lg.WithError(err).Error("GetCartItemByID not found")
		return errHandler.ErrRedirect
	}

	err = cs.CartItemRepo.DeleteCartItem(&item)

	cs.lg.WithError(err).Error("DeleteCartItem")
	return nil
}

func (cs cartService) GetCartData(sessionID string) (items []map[string]interface{}) {
	cartEntity, err := cs.CartEntityRepo.GetCartEntityByStatusAndSessionID(entity.CartOpen, sessionID)
	if err != nil {
		cs.lg.WithError(err).Error("GetCartEntityByStatus not found")
		return
	}

	cartItems, err := cs.CartItemRepo.GetCartItemsByCartID(cartEntity.ID)
	if err != nil {
		cs.lg.WithError(err).Error("GetCartItemByID not found")
		return
	}

	for _, cartItem := range cartItems {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}
	return items
}
