package routers

import (
	"challange-2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(port string) *gin.Engine {
	r := gin.Default()

	r.POST("/orders", controller.CreateOrder)
	r.GET("/orders", controller.GetAllOrders)
	r.PUT("/orders/:orderId", controller.UpdateOrder)
	r.DELETE("/orders/:orderId", controller.DeleteOrder)

	return r
}
