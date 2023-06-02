package routes

import (
	"github.com/Mudassir-Munir/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incommingRoutes.GET("/ordersItems/:orderItem_id", controllers.GetOrderItem())
	//incommingRoutes.GET("/ordersItems/:order_id", controllers.GetOrderItemsByOrder())
	incommingRoutes.POST("/ordersItems", controllers.CreateOrderItem())
	incommingRoutes.PATCH("/ordersItems/:orderItem_id", controllers.UpdateOrderItem())
}
