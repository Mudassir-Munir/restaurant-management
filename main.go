package main

import (
	"os"

	"github.com/Mudassir-Munir/restaurant-management/middleware"
	"github.com/Mudassir-Munir/restaurant-management/routes"
	"github.com/Mudassir-Munir/restaurant-management/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = storage.OpenCollection(storage.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":" + port)

}
