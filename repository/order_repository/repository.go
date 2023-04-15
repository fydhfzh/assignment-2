package order_repository

import (
	"github.com/fydhfzh/assignment-2/entity"
	"github.com/fydhfzh/assignment-2/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(entity.OrderItem) (*entity.OrderItem, errs.ErrMessage)
	GetOrder(int) (*entity.OrderItem, errs.ErrMessage)
	UpdateOrder(int, entity.OrderItem) (*entity.OrderItem, errs.ErrMessage)
	DeleteOrder(int) (errs.ErrMessage)
}