package repository

import (
	"github.com/Arismonx/shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, Logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, Logger}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Error("Failed to list item: %s", err.Error())
		return nil, err
	}

	return itemList, nil
}
