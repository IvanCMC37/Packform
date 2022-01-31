package models

type Customer struct {
	USER_ID      string `json:"user_id" gorm:"primary_key"`
	LOGIN        string `json:"login"`
	PASSWORD     string `json:"password"`
	NAME         string `json:"name"`
	CREDIT_CARDS string `json:"credit_cards"`
	COMPANY_ID   int    `json:"company_id"`
}
