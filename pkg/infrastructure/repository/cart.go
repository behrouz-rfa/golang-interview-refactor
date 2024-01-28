package repository

import (
	"gorm.io/gorm"
	"interview/pkg/core/entity"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (u Repository) CreateCartEntity(cartEntity *entity.Cart) error {
	return u.db.Create(cartEntity).Error
}

func (u Repository) GetCartEntityByStatusAndSessionID(status, sessionID string) (cartEntity entity.Cart, err error) {
	err = u.db.Where("status = ? AND session_id = ?", status, sessionID).First(&cartEntity).Error
	return
}
