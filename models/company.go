package models

type Company struct {
	COMPANY_ID   int    `json:"id" gorm:"primary_key"`
	COMPANY_NAME string `json:"name"`
}
