package routes

import (
	"github.com/Mudassir-Munir/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orders", controllers.GetOrders())
	incommingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incommingRoutes.GET("/orders", controllers.CreateOrder())
	incommingRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
