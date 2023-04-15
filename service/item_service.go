package service

import (
	"fmt"

	"github.com/fydhfzh/assignment-2/entity"
	"github.com/fydhfzh/assignment-2/repository/item_repository"
)

type itemService struct {
	itemRepo item_repository.ItemRepository
}

type ItemService interface {
	FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, error)
}

func NewItemService(itemRepo item_repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (i *itemService) FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, error) {
	items, err := i.itemRepo.FindItemsByItemCodes(itemCodes)

	if err != nil {
		return nil, err
	}

	for _, eachItemCode := range itemCodes {
		isFound := false

		for _, eachItem := range items {
			if eachItemCode == eachItem.ItemCode {
				isFound = true
				break
			}
		}

		if !isFound {
			return nil, fmt.Errorf(fmt.Sprintf("item with code %s does not exist", eachItemCode))
		}
	}

	return items, err
}
