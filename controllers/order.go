package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mudassir-Munir/restaurant-management/models"
	"github.com/Mudassir-Munir/restaurant-management/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = storage.OpenCollection(storage.Client, "food")
var tableCollection *mongo.Collection = storage.OpenCollection(storage.Client, "table")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(ctx, bson.M{})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing all orders items"})
			return
		}

		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allOrders)

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Order

		orderId := c.Param("order_id")

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get order by id"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Order
		var table models.Table

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.E{"table_id", order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "table not found"})
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, err := orderCollection.InsertOne(ctx, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order item"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Order
		var table models.Table

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var updateObj primitive.D

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "table not found"})
				return
			}

			updateObj = append(updateObj, bson.E{"table_id", order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "update order failed"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

// func OrderItemOrderCreater(order models.Order) (string, error) {
// 	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

// 	order.ID = primitive.NewObjectID()
// 	order.Order_id = order.ID.Hex()

// 	_, err := orderCollection.InsertOne(ctx, order)
// 	if err != nil {
// 		//c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order item"})
// 		return "", err
// 	}

// 	return order.Order_id, nil

// }
