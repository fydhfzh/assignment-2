package service

import (
	"net/http"

	"github.com/fydhfzh/assignment-2/dto"
	"github.com/fydhfzh/assignment-2/entity"
	"github.com/fydhfzh/assignment-2/pkg/errs"
	"github.com/fydhfzh/assignment-2/repository/order_repository"
)

type orderService struct {
	orderRepo order_repository.OrderRepository
}

type OrderService interface {
	CreateOrder(dto.NewOrderRequest) (*dto.NewOrderResponse, errs.ErrMessage)
	GetOrder(int) (*dto.GetOrderResponse, errs.ErrMessage)
	UpdateOrder(int, dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, errs.ErrMessage)
	DeleteOrder(int) (errs.ErrMessage)
}

func NewOrderService(orderRepo order_repository.OrderRepository) (OrderService) {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrder(orderPayload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.ErrMessage) {
	var items []entity.Item

	for _, i := range orderPayload.Items {
		var item entity.Item
		item.ItemCode = i.ItemCode
		item.Description = i.Description
		item.Quantity = i.Quantity
		items = append(items, item)
	}

	orderEntity := entity.Order{
		OrderedAt: orderPayload.OrderedAt,
		CustomerName: orderPayload.CustomerName,
	}

	orderItemAggregate := entity.OrderItem{
		OrderData: orderEntity,
		Items: items,
	}

	newOrderItem, err := o.orderRepo.CreateOrder(orderItemAggregate)

	if err != nil {
		return nil, err
	}

	response := &dto.NewOrderResponse{
		StatusCode: http.StatusCreated,
		Data: dto.OrderResponse{
			OrderId: newOrderItem.OrderData.OrderId,
			CustomerName: newOrderItem.OrderData.CustomerName,
		},
	}

	return response, nil
}

func (o *orderService) GetOrder(orderId int) (*dto.GetOrderResponse, errs.ErrMessage) {
	
	orderItemAggregate, err := o.orderRepo.GetOrder(orderId)

	if err != nil {
		return nil, err
	}

	response := &dto.GetOrderResponse{
		Data: dto.OrderResponse{
			OrderId: orderItemAggregate.OrderData.OrderId,
			CreatedAt: orderItemAggregate.OrderData.CreatedAt,
			UpdatedAt: orderItemAggregate.OrderData.UpdatedAt,
			CustomerName: orderItemAggregate.OrderData.CustomerName,
			Items: orderItemAggregate.Items,
		},
	}

	return response, nil
}

func (o *orderService) UpdateOrder(orderId int, orderItemsUpdateRequest dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, errs.ErrMessage) {
	var items []entity.Item

	for _, i := range orderItemsUpdateRequest.Items {
		var item entity.Item
		item.ItemCode = i.ItemCode
		item.Description = i.Description
		item.Quantity = i.Quantity 

		items = append(items, item)
	}

	order := entity.Order{
		OrderedAt: orderItemsUpdateRequest.OrderedAt,
		CustomerName: orderItemsUpdateRequest.CustomerName,  
	}

	orderItemsUpdateAggregate := entity.OrderItem{
		OrderData: order,
		Items: items,
	}

	updatedOrderItemsAggregate, err := o.orderRepo.UpdateOrder(orderId, orderItemsUpdateAggregate)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateOrderResponse{
		StatusCode: http.StatusOK,
		Data: dto.OrderResponse{
			OrderId: updatedOrderItemsAggregate.OrderData.OrderId,
			CreatedAt: updatedOrderItemsAggregate.OrderData.CreatedAt,
			UpdatedAt: updatedOrderItemsAggregate.OrderData.UpdatedAt,
			CustomerName: updatedOrderItemsAggregate.OrderData.CustomerName,
			Items: updatedOrderItemsAggregate.Items,
		},
	}

	return &response, nil
}

func (o *orderService) DeleteOrder(orderId int) (errs.ErrMessage) {

	err := o.orderRepo.DeleteOrder(orderId)

	if err != nil {
		return err
	}

	return nil
}