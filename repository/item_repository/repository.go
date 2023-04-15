package item_repository

import "github.com/fydhfzh/assignment-2/entity"

type ItemRepository interface {
	FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, error)
}
