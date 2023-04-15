package handler

import (
	"github.com/fydhfzh/assignment-2/database"
	"github.com/fydhfzh/assignment-2/repository/order_repository/order_pg"
	"github.com/fydhfzh/assignment-2/service"
	"github.com/gin-gonic/gin"
)

const PORT = ":3000"

func StartApp() {
	database.InitializeDatabase()
	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewOrderPG(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := NewOrderHandler(orderService)

	router := gin.Default()
	router.GET("/orders/:orderId", orderHandler.GetOrder)
	router.PUT("/orders/:orderId", orderHandler.UpdateOrder)
	router.POST("/orders", orderHandler.CreateOrder)
	router.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	router.Run(PORT)
}