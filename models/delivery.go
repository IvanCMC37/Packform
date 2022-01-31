package models

type Delivery struct {
	ID                 int `json:"id" gorm:"primary_key"`
	ORDER_ITEM_ID      int `json:"order_item_id" gorm:"foreignkey:ID"`
	DELIVERED_QUANTITY int `json:"delivered_quantity"`
}
