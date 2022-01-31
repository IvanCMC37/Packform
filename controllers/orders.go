package controllers

import (
	"net/http"
	"packform/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	// type Result struct {
	// 	ORDER_NAME   string
	// 	COMPANY_NAME string
	// 	NAME         string
	// 	ORDER_DATE   string
	// }

	// var results []Result

	// models.DB.Table("companies").Select("orders.ORDER_NAME, companies.COMPANY_NAME, customers.NAME, orders.ORDER_DATE").
	// 	Joins("INNER join customers on companies.COMPANY_ID = customers.COMPANY_ID").
	// 	Joins("INNER join orders on orders.CUSTOMER_ID = customers.USER_ID").Scan(&results)

	// type Result struct {
	// 	ORDER_ID int
	// 	TOTAL    float64
	// 	PRODUCTS string
	// }
	// var results []Result

	// models.DB.Table("order_items").Select("order_id,sum(quantity*price_per_unit) as TOTAL, string_agg(product,';') as PRODUCTS").Group("order_id").Order("order_id").Scan(&results)

	type Result struct {
		ORDER_ITEM_ID      int
		DELIVERED_QUANTITY int
		DELIVERED_AMOUNT   float64
	}
	var results []Result

	models.DB.Table("order_items").Select("order_item_id, sum(delivered_quantity) as DELIVERED_QUANTITY, sum(DELIVERED_QUANTITY*price_per_unit) as DELIVERED_AMOUNT").
		Joins(("INNER join deliveries on deliveries.order_item_id = order_items.id")).Group("order_item_id").Order("order_item_id").Scan(&results)

	// models.DB.Raw("select order_item_id,sum(delivered_quantity) as DELIVERED_QUANTITY from deliveries group by order_item_id").Scan(&results)
	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})

}
