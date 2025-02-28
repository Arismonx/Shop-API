package service

import (
	_itemShopModel "github.com/Arismonx/shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing() ([]*_itemShopModel.Item, error)
}
