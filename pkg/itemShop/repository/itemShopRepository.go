package repository

import "github.com/Arismonx/shop-api/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
