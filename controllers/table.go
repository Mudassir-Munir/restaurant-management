package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := tableCollection.Find(ctx, bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get list of table"})
			return
		}

		var allTables []bson.M
		if err := result.All(ctx, &allTables); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allTables)
	}
}

func GetTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func CreateTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
