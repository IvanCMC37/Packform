package models

import (
	"time"
)

type Order struct {
	ID          int       `json:"id" gorm:"primary_key"`
	CREATED_AT  time.Time `json:"CREATED_AT"`
	ORDER_NAME  string    `json:"ORDER_NAME"`
	CUSTOMER_ID string    `json:"CUSTOMER_ID" gorm:"foreignkey:USER_ID"`
}
