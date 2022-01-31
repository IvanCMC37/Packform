package models

type Order struct {
	ID          int    `json:"id" gorm:"primary_key"`
	ORDER_DATE  string `json:"order_date"`
	ORDER_NAME  string `json:"order_name"`
	CUSTOMER_ID string `json:"customer_id" gorm:"foreignkey:USER_ID"`
}
