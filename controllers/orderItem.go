package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// implement me
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// implement me
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	// implement me
	return nil, nil
}

func GetOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	// implement me
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	// implement me
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
     // implement me
	}
}
