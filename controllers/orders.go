package controllers

import (
	"net/http"
	"packform/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	type Result struct {
		ORDER_NAME         string
		COMPANY_NAME       string
		NAME               string
		ORDER_DATE         string
		PRODUCTS           string
		TOTAL              float64
		ORDER_QUANTITY     int
		DELIVERED_QUANTITY int
		DELIVERED_AMOUNT   float64
	}

	var results []Result

	models.DB.Table("companies").Select("ORDER_NAME, COMPANY_NAME, NAME, ORDER_DATE, LOWER(string_agg(product,';')) as PRODUCTS, sum(DISTINCT(quantity*price_per_unit)) as TOTAL" +
		",sum(DISTINCT(quantity)) as ORDER_QUANTITY, sum(delivered_quantity) as DELIVERED_QUANTITY, sum(DELIVERED_QUANTITY*price_per_unit) as DELIVERED_AMOUNT").
		Joins("INNER join customers on companies.COMPANY_ID = customers.COMPANY_ID").
		Joins("INNER join orders on orders.CUSTOMER_ID = customers.USER_ID").
		Joins("INNER join order_items on order_items.order_id = orders.id").
		Joins("INNER join deliveries as d on d.order_item_id = order_items.id").
		Group("order_items.order_id, ORDER_NAME,COMPANY_NAME,NAME, ORDER_DATE").Scan(&results)

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})

}
