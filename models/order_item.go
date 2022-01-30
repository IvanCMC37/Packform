package models

type Order_item struct {
	ID             int     `json:"id" gorm:"primary_key"`
	ORDER_ID       int     `json:"order_id"`
	PRICE_PER_UNIT float64 `json:"price_per_unit"`
	QUANTITY       int     `json:"quantity"`
	PRODUCT        string  `json:"product"`
}
