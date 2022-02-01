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

	models.DB.Table("companies").Select("ORDER_ID,ORDER_NAME, COMPANY_NAME, NAME, string_agg(DISTINCT(ORDER_DATE),';') as ORDER_DATE, string_agg(DISTINCT(product),';')" +
		" as PRODUCTS, sum(DISTINCT(TOTAL)) as TOTAL,sum(DISTINCT(quantity)) as ORDER_QUANTITY, sum(delivered_quantity) as DELIVERED_QUANTITY, sum(DELIVERED_AMOUNT) as DELIVERED_AMOUNT").
		Joins("INNER join customers on companies.COMPANY_ID = customers.COMPANY_ID").
		Joins("INNER JOIN(select o.id as ORDER_ID, string_agg(DISTINCT(order_date),';') as ORDER_DATE, order_name, customer_id, sum(DISTINCT(quantity*price_per_unit)) as total" +
			",LOWER(string_agg(distinct(product),';')) as PRODUCT, sum(DISTINCT(quantity)) as quantity" +
			",sum(coalesce(delivered_quantity,0)) as DELIVERED_QUANTITY" +
			",sum(coalesce(delivered_quantity,0)*price_per_unit) as DELIVERED_AMOUNT" +
			" from orders as o" +
			" join order_items as oi on oi.order_id=o.id" +
			" left join deliveries as d on d.order_item_id = oi.id" +
			" group by o.id,order_name,customer_id,oi.id,d.order_item_id) as foo on foo.CUSTOMER_ID = customers.USER_ID").
		Group("ORDER_ID, ORDER_NAME,COMPANY_NAME,NAME").
		Order("ORDER_ID").Scan(&results)

	c.JSON(http.StatusOK, gin.H{
		"orders": results,
	})

}
