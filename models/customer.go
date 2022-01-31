// package models

// type Customer struct {
// 	USER_ID      string `json:"USER_ID" gorm:"primary_key"`
// 	LOGIN        string `json:"LOGIN"`
// 	PASSWORD     string `json:"PASSWORD"`
// 	NAME         string `json:"NAME"`
// 	CREDIT_CARDS string `json:"CREDIT_CARDS"`
// 	COMPANY_ID   int    `json:"COMPANY_ID"`
// }
package models

type Customer struct {
	USER_ID      string `json:"user_id" gorm:"primary_key"`
	LOGIN        string `json:"login"`
	PASSWORD     string `json:"password"`
	NAME         string `json:"name"`
	CREDIT_CARDS string `json:"credit_cards"`
	COMPANY_ID   int    `json:"company_id"`
}
