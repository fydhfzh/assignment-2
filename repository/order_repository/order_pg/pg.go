package order_pg

import (
	"database/sql"
	"fmt"

	"github.com/fydhfzh/assignment-2/entity"
	"github.com/fydhfzh/assignment-2/pkg/errs"
	"github.com/fydhfzh/assignment-2/repository/order_repository"
)

type orderPG struct {
	db *sql.DB
}

func NewOrderPG(db *sql.DB) (order_repository.OrderRepository) {
	return &orderPG{
		db: db,
	}
}

const (
	createOrderQuery = `
		INSERT INTO "orders" (customer_name, ordered_at)
		VALUES ($1, $2)
		RETURNING order_id, customer_name, ordered_at;
	`

	createItemQuery = `
		INSERT INTO "items" (item_code, description, quantity, order_id)
		VALUES ($1, $2, $3, $4)
		RETURNING item_id, created_at, updated_at, item_code, description, quantity, order_id
	`

	getOrderQuery = `
		SELECT created_at, updated_at, customer_name
		FROM "orders" WHERE order_id = $1;
	`

	getItemsQuery = `
		SELECT item_id, created_at, updated_at, item_code, description, quantity, order_id
		FROM "items" WHERE order_id = $1; 
	`

	updateItemsQuery = `
		UPDATE "items"
		SET description = $2, quantity = $3, updated_at = now()
		WHERE order_id = $4 AND item_code = $1
		RETURNING item_id, created_at, updated_at, item_code, description, quantity, order_id;
	`

	updateOrderQuery = `
		UPDATE "orders"
		SET ordered_at = $1, customer_name = $2, updated_at = now()
		WHERE order_id = $3
		RETURNING order_id, created_at, updated_at, customer_name;
	`

	deleteOrderQuery = `
		DELETE FROM "orders"
		WHERE order_id = $1;
	`

	deleteItemsQuery = `
		DELETE FROM "items"
		WHERE order_id = $1;
	`
)

func (o *orderPG) CreateOrder(orderItemAggregate entity.OrderItem) (*entity.OrderItem, errs.ErrMessage){
	tx, err := o.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("internal server error")
	}

	row := tx.QueryRow(createOrderQuery, orderItemAggregate.OrderData.CustomerName, orderItemAggregate.OrderData.OrderedAt)

	var createdOrder entity.Order

	err = row.Scan(&createdOrder.OrderId, &createdOrder.CustomerName, &createdOrder.OrderedAt)

	if err != nil {
		return nil, errs.NewUnprocessableEntityError("invalid request body")
	}

	var createdItems []entity.Item

	for _, item := range orderItemAggregate.Items {
		row = tx.QueryRow(item.ItemCode, item.Description, item.Quantity, orderItemAggregate.OrderData.OrderId)
		
		var item entity.Item

		err := row.Scan(&item.ItemId, &item.CreatedAt, &item.UpdatedAt, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)

		if err != nil {
			tx.Rollback()
			return nil, errs.NewUnprocessableEntityError("invalid request body")	
		}

		createdItems = append(createdItems, item)
	}

	err = tx.Commit()

	if err != nil {
		return nil, errs.NewInternalServerError("internal server error")
	}

	createdOrderItemAggregate := entity.OrderItem{
		OrderData: createdOrder,
		Items: createdItems,
	}

	return &createdOrderItemAggregate, nil
}

func (o *orderPG) GetOrder(orderId int) (*entity.OrderItem, errs.ErrMessage) {
	row := o.db.QueryRow(getOrderQuery, orderId)

	var order entity.Order
	order.OrderId = orderId

	err := row.Scan(&order.CreatedAt, &order.UpdatedAt, &order.CustomerName)

	if err != nil {
		return nil, errs.NewUnprocessableEntityError("invalid request body")
	}

	rows, err := o.db.Query(getItemsQuery, orderId)

	if err != nil {
		return nil, errs.NewUnprocessableEntityError("invalid request body")
	}

	var items []entity.Item

	for rows.Next() {
		var item entity.Item
		err = rows.Scan(&item.ItemId, &item.CreatedAt, &item.UpdatedAt, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
		if err != nil {
			return nil, errs.NewUnprocessableEntityError("invalid request body")
		}

		items = append(items, item)
	}

	orderItemAggregate := entity.OrderItem{
		OrderData: order,
		Items: items,
	}

	return &orderItemAggregate, nil
}

func (o *orderPG) UpdateOrder(orderId int, orderItemUpdate entity.OrderItem) (*entity.OrderItem, errs.ErrMessage) {
	tx, err := o.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("internal server error")
	}

	var updatedItems []entity.Item

	for _, item := range orderItemUpdate.Items {
		row := tx.QueryRow(updateItemsQuery, item.ItemCode, item.Description, item.Quantity, orderId)
		
		var updatedItem entity.Item

		err := row.Scan(&updatedItem.ItemId, &updatedItem.CreatedAt, &updatedItem.UpdatedAt, &updatedItem.ItemCode, &updatedItem.Description, &updatedItem.Quantity, &updatedItem.OrderId)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil, errs.NewUnprocessableEntityError("invalid request body")
		}

		updatedItems = append(updatedItems, updatedItem)
	}

	var updatedOrder entity.Order

	row := tx.QueryRow(updateOrderQuery, orderItemUpdate.OrderData.OrderedAt, orderItemUpdate.OrderData.CustomerName, orderId)

	err = row.Scan(&updatedOrder.OrderId, &updatedOrder.CreatedAt, &updatedOrder.UpdatedAt, &updatedOrder.CustomerName)

	if err != nil {
		tx.Rollback()
		fmt.Printf("error: %v\n", err)
		return nil, errs.NewUnprocessableEntityError("invalid request body")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		fmt.Printf("error: %v\n", err)
		return nil, errs.NewInternalServerError("internal server error")
	}

	updatedOrderItemAggregate := entity.OrderItem{
		OrderData: updatedOrder,
		Items: updatedItems,
	}

	return &updatedOrderItemAggregate, nil
}

func (o *orderPG) DeleteOrder(orderId int) (errs.ErrMessage) {
	tx, err := o.db.Begin()

	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return errs.NewInternalServerError("internal server error")
	}

	_, err = tx.Exec(deleteItemsQuery, orderId)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return errs.NewUnprocessableEntityError("invalid request body")
	}

	_, err = tx.Exec(deleteOrderQuery, orderId)

	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return errs.NewUnprocessableEntityError("invalid request body")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return errs.NewInternalServerError("internal server error")
	}

	return nil
}