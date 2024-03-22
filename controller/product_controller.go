package controller

import (
	"challange-2/database"
	"challange-2/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		handleError(ctx, err, "Failed to bind JSON for creating order")
		return
	}

	if err := db.Create(&order).Error; err != nil {
		handleError(ctx, err, "Failed to create order")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"msg":     "Order created successfully",
	})
}

func GetAllOrders(ctx *gin.Context) {
	db := database.GetDB()
	orders := []models.Order{}

	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		handleError(ctx, err, "Failed to fetch orders from the database")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  orders,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}
	success := true
	msg := ""

	if err := ctx.ShouldBindJSON(&order); err != nil {
		handleError(ctx, err, "Failed to bind JSON for updating order")
		return
	}

	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		handleError(ctx, err, "Invalid order ID")
		return
	}
	order.Order_ID = uint(orderID)

	for index := range order.Items {
		if order.Items[index].Item_ID == 0 {
			success = false
			msg = "Item ID can't be empty"
			break
		}
		order.Items[index].OrderID = order.Order_ID
	}

	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": success,
			"msg":     msg,
		})
		return
	}

	if err := db.Save(&order).Error; err != nil {
		handleError(ctx, err, "Failed to update order in the database")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		handleError(ctx, err, "Invalid order ID")
		return
	}

	if err := db.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		handleError(ctx, err, "Failed to delete items associated with the order")
		return
	}

	if err := db.Where("order_id = ?", orderID).Delete(&models.Order{}).Error; err != nil {
		handleError(ctx, err, "Failed to delete order from the database")
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"msg":     "Order has been successfully deleted",
	})
}

func handleError(ctx *gin.Context, err error, message string) {
	log.Println(err)
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   err.Error(),
		"msg":     message,
	})
}
