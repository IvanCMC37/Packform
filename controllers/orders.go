package controllers

import (
	"net/http"
	"packform/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	type Result struct {
		ORDER_NAME   string
		COMPANY_NAME string
		NAME         string
		ORDER_DATE   time.Time
	}

	var results []Result

	models.DB.Table("companies").Select("orders.ORDER_NAME, companies.COMPANY_NAME, customers.NAME, orders.CREATED_AT").
		Joins("INNER join customers on companies.COMPANY_ID = customers.COMPANY_ID").
		Joins("INNER join orders on orders.CUSTOMER_ID = customers.USER_ID").Scan(&results)

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
	// var order_items []models.Order_item
	// models.DB.Where("order_id = ?", "1").Find(&order_items)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": order_items,
	// })
}
