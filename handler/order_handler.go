package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fydhfzh/assignment-2/dto"
	"github.com/fydhfzh/assignment-2/service"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.OrderService
}

type OrderHandler interface {
	CreateOrder(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

func NewOrderHandler(orderService service.OrderService) (*orderHandler) {
	return &orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) CreateOrder(ctx *gin.Context){
	var newOrderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errorMessage": "invalid request body",
		})

		return
	}

	fmt.Printf("%v", newOrderRequest)

	response, err := o.orderService.CreateOrder(newOrderRequest)

	if err == sql.ErrTxDone {
		fmt.Printf("sql: transaction already have been committed or rollback")
		ctx.JSON(err.Status(), gin.H{
			"errorMessage": err.Message(),
		})

		return
	} else if err == sql.ErrNoRows {
		fmt.Printf("sql: item code not found")
		ctx.JSON(err.Status(), gin.H{
			"errorMessage": err.Message(),
		})

		return
	} else if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		ctx.JSON(err.Status(), gin.H{
			"errorMessage": err.Message(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (o *orderHandler) GetOrder(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("orderId"))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errorMessage": "page not found",
		})

		return
	}

	response, responseErr := o.orderService.GetOrder(id)

	if responseErr == sql.ErrNoRows {
		ctx.JSON(responseErr.Status(), gin.H{
			"errorMessage": responseErr.Message(),
		})

		return
	}else if responseErr != nil {
		fmt.Printf("error: %v\n", err)
		ctx.JSON(responseErr.Status(), gin.H{
			"errorMessage": responseErr.Message(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (o *orderHandler) UpdateOrder(ctx *gin.Context){
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "page not found",
		})

		return
	}

	var orderRequest dto.UpdateOrderRequest

	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "invalid request body",
		})

		return
	}

	response, responseErr := o.orderService.UpdateOrder(orderId, orderRequest)
	
	if responseErr != nil {
		ctx.JSON(responseErr.Status(), gin.H{
			"errorMessage": responseErr.Message(),
		})
		
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (o *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %d Not Found", orderId),
		})

		return
	}

	responseErr := o.orderService.DeleteOrder(orderId)

	if err != nil {
		ctx.JSON(responseErr.Status(), gin.H{
			"errorMessage": responseErr.Message(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "order data deleted successfully",
	})
}