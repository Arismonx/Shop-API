package service

import (
	_itemShopModel "github.com/Arismonx/shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Arismonx/shop-api/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing() ([]*_itemShopModel.Item, error) {
	itemList, err := s.itemShopRepository.Listing()
	if err != nil {
		return nil, err
	}

	itemShopModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemList {
		itemShopModelList = append(itemShopModelList, item.ToItemModel())
	}

	return itemShopModelList, nil
}
