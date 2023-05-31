package routes

import (
	"github.com/Mudassir-Munir/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/menus", controllers.GetMenus())
	incommingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incommingRoutes.GET("/menus", controllers.CreateMenu())
	incommingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
