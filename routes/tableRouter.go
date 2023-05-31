package routes

import (
	"github.com/Mudassir-Munir/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)


func TableRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/tables", controllers.GetTables())
	incommingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incommingRoutes.GET("/tables", controllers.CreateTable())
	incommingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
