package models

type Customer struct {
	USER_ID      string `json:"USER_ID" gorm:"primary_key"`
	LOGIN        string `json:"LOGIN"`
	PASSWORD     string `json:"PASSWORD"`
	NAME         string `json:"NAME"`
	CREDIT_CARDS string `json:"CREDIT_CARDS"`
	COMPANY_ID   int    `json:"COMPANY_ID"`
}
